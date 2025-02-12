# 📈 Stock Market Aggregator

This repository contains a **Stock Market Aggregator** script built with Go.

## 🧑‍💻 About the Project

The project is designed to collect **stock market data** from free APIs and aggregate the data to show the **average price** of stocks.

### 🎯 Project Goal

The main goal of this project is to implement **Go-native concurrency patterns** and apply them in a real-world scenario of fetching and processing stock data. The system fetches data from multiple APIs in parallel, processes it, and aggregates the data for further analysis.

### 📋 Project Breakdown

**Main Program** initializes the application, sets up logging, and controls the flow between data fetching and aggregation.

- Loads environment variables (using the godotenv package).
- Initializes logging for information and error messages.
- Creates instances of the Fetcher and Aggregator modules.
- Fetches stock data by calling the Fetcher's Fetch function.
- Passes the fetched data to the Aggregator to aggregate the stock prices.

**Data Fetching** retrieves stock data from free stock market APIs, concurrently.

- Fetches stock data using multiple free APIs (such as FMP Cloud or Alpha Vantage).
- Each URL request is handled concurrently using Go's goroutines, ensuring efficient data fetching.
- Errors are logged if there are any issues with the API requests.
- The Fetch method of the Fetcher returns the raw data for further processing.

**Data Aggregation** aggregates the fetched stock data.

- Receives the fetched stock data from the Fetcher.
- For now, the method only logs that aggregation is attempted (to be expanded in future development).
