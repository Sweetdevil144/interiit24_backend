INSERT INTO companies ("name", "country", "country_code", "market_cap", "diversity", "created_at", "updated_at")
VALUES
('Zooxo', 'Ukraine', 'UA', 2340000000, 43.5, NOW(), NOW()),
('Gabtype', 'Brazil', 'BR', 30670000000, 53.4, NOW(), NOW()),
('Quinu', 'China', 'CN', 26320000, 83, NOW(), NOW()),
('Youbridge', 'Colombia', 'CO', 695780000, 89.2, NOW(), NOW()),
('Oyope', 'Ukraine', 'UA', 1280000000, 23.3, NOW(), NOW()),
('Rhynyx', 'Indonesia', 'ID', 8390000000, 84.6, NOW(), NOW()),
('Babblestorm', 'Indonesia', 'ID', 9260000000, 48.1, NOW(), NOW()),
('Jabberbean', 'Croatia', 'HR', 113200000, 62.1, NOW(), NOW()),
('Kamba', 'Ivory Coast', 'CI', 121900000, 48.4, NOW(), NOW()),
('Abatz', 'Indonesia', 'ID', 488160000, 74.7, NOW(), NOW());


INSERT INTO financial_data ("company_id", "year", "stock_price", "expense", "revenue", "market_share", "created_at", "updated_at")
VALUES
((SELECT id FROM companies WHERE name = 'Zooxo'), 2015, 636190000, 20101562.63, 11889859.91, 28.24, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Zooxo'), 2016, 36170000000, 51007842.36, 12926239.97, 42.01, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Zooxo'), 2017, NULL, 62056576.98, 26106119.10, 73.14, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Zooxo'), 2018, 1060000000, 79478124.95, 86658338.25, 90.64, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Zooxo'), 2019, 308220000, 52885481.03, 21961920.16, 56.8, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Zooxo'), 2020, 514130000, 71228506.48, 93825082.89, 48.79, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Zooxo'), 2021, 21660000, 17224689.73, 50393952.13, 35.4, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Zooxo'), 2022, 555700000, 96688655.13, 78632245.86, 65.88, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Zooxo'), 2023, 355280000, 67582493.42, 95627554.77, 38.85, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Zooxo'), 2024, NULL, 81035180.73, 39485065.72, 77.69, NOW(), NOW());

INSERT INTO financial_data ("company_id", "year", "stock_price", "expense", "revenue", "market_share", "created_at", "updated_at")
VALUES
((SELECT id FROM companies WHERE name = 'Gabtype'), 2015, 1330000000, 4900246.33, 19402078.01, 75.06, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Gabtype'), 2016, 127360000000, 11650445.38, 73943317.58, 71.2, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Gabtype'), 2017, 1970000, 14571356.05, 12844161.78, 11.7, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Gabtype'), 2018, 192010000, 1538107.74, 32226487.57, 37.98, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Gabtype'), 2019, 213580000, 61686123.87, 49520814.50, 67.95, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Gabtype'), 2020, 36760000000, 58635750.09, 83288693.79, 40.41, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Gabtype'), 2021, 29720000, 21049564.20, 62086303.35, 65.85, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Gabtype'), 2022, 61780000, 29790519.29, 79287568.31, 55.78, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Gabtype'), 2023, NULL, 32879281.41, 81877647.40, 96.02, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Gabtype'), 2024, 28430000000, 64622482.02, 8402729.73, 4.96, NOW(), NOW());


INSERT INTO "financial_data" ("company_id", "year", "stock_price", "expense", "revenue", "market_share", "created_at", "updated_at")
VALUES
((SELECT id FROM companies WHERE name = 'Youbridge'), 2015, 750000000, 18000000.00, 70000000.00, 25.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Youbridge'), 2016, 800000000, 20000000.00, 75000000.00, 30.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Youbridge'), 2017, 850000000, 22000000.00, 80000000.00, 35.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Youbridge'), 2018, 900000000, 25000000.00, 85000000.00, 40.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Youbridge'), 2019, 950000000, 27000000.00, 90000000.00, 45.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Youbridge'), 2020, 1000000000, 30000000.00, 95000000.00, 50.00, NOW(), NOW());

INSERT INTO "financial_data" ("company_id", "year", "stock_price", "expense", "revenue", "market_share", "created_at", "updated_at")
VALUES
((SELECT id FROM companies WHERE name = 'Quinu'), 2015, 20000000, 3000000.00, 12000000.00, 15.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Quinu'), 2016, 25000000, 4000000.00, 15000000.00, 20.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Quinu'), 2017, 30000000, 5000000.00, 18000000.00, 25.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Quinu'), 2018, 35000000, 6000000.00, 20000000.00, 30.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Quinu'), 2019, 40000000, 7000000.00, 22000000.00, 35.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Quinu'), 2020, 45000000, 8000000.00, 25000000.00, 40.00, NOW(), NOW());

INSERT INTO "financial_data" ("company_id", "year", "stock_price", "expense", "revenue", "market_share", "created_at", "updated_at")
VALUES
((SELECT id FROM companies WHERE name = 'Rhynyx'), 2015, 7500000000, 15000000.00, 70000000.00, 55.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Rhynyx'), 2016, 8000000000, 20000000.00, 75000000.00, 60.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Rhynyx'), 2017, 8500000000, 22000000.00, 80000000.00, 65.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Rhynyx'), 2018, 9000000000, 25000000.00, 85000000.00, 70.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Rhynyx'), 2019, 9500000000, 27000000.00, 90000000.00, 75.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Rhynyx'), 2020, 10000000000, 30000000.00, 95000000.00, 80.00, NOW(), NOW());

INSERT INTO "financial_data" ("company_id", "year", "stock_price", "expense", "revenue", "market_share", "created_at", "updated_at")
VALUES
((SELECT id FROM companies WHERE name = 'Babblestorm'), 2015, 9000000000, 18000000.00, 85000000.00, 40.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Babblestorm'), 2016, 9500000000, 20000000.00, 90000000.00, 45.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Babblestorm'), 2017, 10000000000, 22000000.00, 95000000.00, 50.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Babblestorm'), 2018, 10500000000, 25000000.00, 100000000.00, 55.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Babblestorm'), 2019, 11000000000, 27000000.00, 105000000.00, 60.00, NOW(), NOW()),
((SELECT id FROM companies WHERE name = 'Babblestorm'), 2020, 11500000000, 30000000.00, 110000000.00, 65.00, NOW(), NOW());
