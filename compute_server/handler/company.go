package handler

import (
	"compute_server/database"
	"compute_server/helpers"
	"compute_server/model"
	"log"
	"math"
	"time"

	"github.com/gofiber/fiber/v2"
)

func SearchCompanies(c *fiber.Ctx) error {
	query := c.Query("name")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Query parameter is required"})
	}
	log.Println(query)
	var companies []model.Company
	if err := database.DB.Where("name ILIKE ? OR country_code ILIKE ? OR country ILIKE ?", "%"+query+"%", "%"+query+"%", "%"+query+"%").Find(&companies).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to fetch companies"})
	}
	return c.JSON(companies)
}

func ComputeData(c *fiber.Ctx) error {
	userId, err := helpers.GetUserFromContext(c)
	companyID, err := c.ParamsInt("companyID")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid company ID"})
	}
	var company model.Company
	if err := database.DB.Preload("Financials").First(&company, companyID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Company not found"})
	}
	result := make(map[string]interface{})
	var countSameCountry int64
	database.DB.Model(&model.Company{}).Where("country = ?", company.Country).Count(&countSameCountry)
	result["same_country_count"] = countSameCountry

	var countGreaterDiversity int64
	database.DB.Model(&model.Company{}).Where("country = ? AND diversity > ?", company.Country, company.Diversity).Count(&countGreaterDiversity)
	result["greater_diversity_count"] = countGreaterDiversity

	changes := make([]map[string]interface{}, len(company.Financials)-1)
	for i := 1; i < len(company.Financials); i++ {
		previous := company.Financials[i-1]
		current := company.Financials[i]
		changes[i-1] = map[string]interface{}{
			"year":                current.Year,
			"stock_price_change":  calculatePercentageChange(previous.StockPrice, current.StockPrice),
			"expense_change":      calculatePercentageChange(previous.Expense, current.Expense),
			"revenue_change":      calculatePercentageChange(previous.Revenue, current.Revenue),
			"market_share_change": calculatePercentageChange(previous.MarketShare, current.MarketShare),
		}
	}
	result["financial_changes"] = changes
	result["greater_metrics_domestic"] = CountGreaterMetricsDomestic(company)
	result["greater_metrics_global"] = CountGreaterMetricsGlobal(company)
	analysisResult, err := AnalyzeCompanyStatistics(company.ID)
	if err != nil {
		analysisResult = nil
	}
	result["analysis"] = analysisResult

	var searchHistory = model.SearchHistory{
		UserID:       userId,
		CompanyID:    uint(companyID),
		StoredResult: result,
		Timestamp:    time.Now(),
	}
	if err := database.DB.Create(&searchHistory).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to store search history"})
	}
	return c.JSON(result)
}

func calculatePercentageChange(previous, current float64) float64 {
	if previous == 0 {
		return 0
	}
	return ((current - previous) / previous) * 100
}

func CountGreaterMetricsDomestic(company model.Company) int64 {
	var count int64
	latestFinancial := company.Financials[len(company.Financials)-1]
	database.DB.Model(&model.Company{}).
		Where("country = ? AND (id != ? AND (id IN (SELECT company_id FROM financial_data WHERE stock_price > ? OR market_share > ? OR revenue > ? OR expense > ?)))",
			company.Country, company.ID, latestFinancial.StockPrice, latestFinancial.MarketShare,
			latestFinancial.Revenue, latestFinancial.Expense).
		Count(&count)

	return count
}

func CountGreaterMetricsGlobal(company model.Company) int64 {
	var count int64
	latestFinancial := company.Financials[len(company.Financials)-1]
	database.DB.Model(&model.Company{}).
		Where("id != ? AND (id IN (SELECT company_id FROM financial_data WHERE stock_price > ? OR market_share > ? OR revenue > ? OR expense > ?))",
			company.ID, latestFinancial.StockPrice, latestFinancial.MarketShare,
			latestFinancial.Revenue, latestFinancial.Expense).
		Count(&count)

	return count
}

func AnalyzeCompanyStatistics(companyID uint) (fiber.Map, error) {
	var company model.Company
	if err := database.DB.Preload("Financials").First(&company, companyID).Error; err != nil {
		return nil, err
	}
	cagr := calculateCAGR(company.Financials)
	volatility := calculateVolatility(company.Financials)
	return fiber.Map{
		"cagr":       cagr,
		"volatility": volatility,
	}, nil
}
func calculateCAGR(financials []model.FinancialData) float64 {
	if len(financials) < 2 {
		return 0
	}
	startValue := financials[0].Revenue
	endValue := financials[len(financials)-1].Revenue
	years := float64(financials[len(financials)-1].Year - financials[0].Year)
	if startValue == 0 {
		return 0
	}
	return (math.Pow(endValue/startValue, 1/years) - 1) * 100
}

func calculateVolatility(financials []model.FinancialData) float64 {
	if len(financials) == 0 {
		return 0
	}
	mean := 0.0
	revenue := make([]float64, len(financials))
	for i, data := range financials {
		revenue[i] = data.Revenue
		mean += data.Revenue
	}
	mean /= float64(len(financials))
	var variance float64
	for _, value := range revenue {
		variance += math.Pow(value-mean, 2)
	}
	variance /= float64(len(financials))
	return math.Sqrt(variance)
}

func FetchFinancialData(c *fiber.Ctx) error {
	companyId := c.Params("companyId")
	var financialData []model.FinancialData
	if err := database.DB.Where("company_id = ?", companyId).Find(&financialData).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to fetch financial data"})
	}
	return c.JSON(financialData)
}
