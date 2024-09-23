package handler

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"server/database"
	"server/model"
	"math"
)

func SearchCompanies(c *fiber.Ctx) error {
	query := c.Query("query")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Query parameter is required"})
	}
	var companies []models.Company
	if err := database.DB.Where("name LIKE ? OR code LIKE ?", "%"+query+"%", "%"+query+"%").Find(&companies).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Unable to fetch companies"})
	}

	return c.JSON(companies)
}

func ComputeData(c *fiber.Ctx) error {
	companyID, err := c.ParamsUInt("companyID") 
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid company ID"})
	}

	var company models.Company
	if err := database.DB.Preload("Financials").First(&company, companyID).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Company not found"})
	}

	result := make(map[string]interface{})

	var countSameCountry int64
	database.DB.Model(&models.Company{}).Where("country = ?", company.Country).Count(&countSameCountry)
	result["same_country_count"] = countSameCountry

	var countGreaterDiversity int64
	database.DB.Model(&models.Company{}).Where("country = ? AND diversity > ?", company.Country, company.Diversity).Count(&countGreaterDiversity)
	result["greater_diversity_count"] = countGreaterDiversity

	changes := make([]map[string]interface{}, len(company.Financials)-1)
	for i := 1; i < len(company.Financials); i++ {
		previous := company.Financials[i-1]
		current := company.Financials[i]
		changes[i-1] = map[string]interface{}{
			"year":                   current.Year,
			"stock_price_change":     calculatePercentageChange(previous.StockPrice, current.StockPrice),
			"expense_change":         calculatePercentageChange(previous.Expense, current.Expense),
			"revenue_change":         calculatePercentageChange(previous.Revenue, current.Revenue),
			"market_share_change":    calculatePercentageChange(previous.MarketShare, current.MarketShare),
		}
	}
	result["financial_changes"] = changes

	result["greater_metrics_domestic"] = CountGreaterMetricsDomestic(company)
	result["greater_metrics_global"] = CountGreaterMetricsGlobal(company)
	analysisResult, err := AnalyzeCompanyStatistics(companyID)
	if err != nil {
		analysisResult = nil
	}
	result["analysis"] = analysisResult
	return c.JSON(result) 
}

func calculatePercentageChange(previous, current float64) float64 {
	if previous == 0 {
		return 0
	}
	return ((current - previous) / previous) * 100
}

func CountGreaterMetricsDomestic(company models.Company) int64 {
	var count int64
	latestFinancial := company.Financials[len(company.Financials)-1]

	database.DB.Model(&models.Company{}).
		Where("country = ? AND (stock_price > ? OR market_share > ? OR revenue > ? OR expense > ?)",
			company.Country, latestFinancial.StockPrice, latestFinancial.MarketShare,
			latestFinancial.Revenue, latestFinancial.Expense).
		Count(&count)
	return count
}

func CountGreaterMetricsGlobal(company models.Company) int64 {
	var count int64
	latestFinancial := company.Financials[len(company.Financials)-1]

	database.DB.Model(&models.Company{}).
		Where("(stock_price > ? OR market_share > ? OR revenue > ? OR expense > ?)",
			latestFinancial.StockPrice, latestFinancial.MarketShare,
			latestFinancial.Revenue, latestFinancial.Expense).
		Count(&count)
	return count
}

func AnalyzeCompanyStatistics(companyID uint) (fiber.Map, error) {
	var company models.Company
	if err := database.DB.Preload("Financials").First(&company, companyID).Error; err != nil {
		return nil, err
	}
	cagr := calculateCAGR(company.Financials)
	volatility := calculateVolatility(company.Financials)
	return fiber.Map{
		"cagr":      cagr,
		"volatility": volatility,
	}, nil
}
func calculateCAGR(financials []models.FinancialData) float64 {
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

func calculateVolatility(financials []models.FinancialData) float64 {
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