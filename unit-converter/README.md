# 🌐 Unit Converter Web App

This is a simple web application that allows users to convert between different units of measurement. It supports conversions for **length**, **weight**, and **temperature**. In their respective folders (with-http and with-gin), we have the development to understand how it is done with the standard library and the advantage of using the framework.

---

## 🚀 Features

- Convert between various units of:
  - **Length**: millimeter, centimeter, meter, kilometer, inch, foot, yard, mile
  - **Weight**: milligram, gram, kilogram, ounce, pound
  - **Temperature**: Celsius, Fahrenheit, Kelvin
- Clean and minimal web interface
- Input a value, choose the units to convert **from** and **to**, and view the result
- No database required — all conversions are handled server-side

---

## 🛠️ How It Works

- The application consists of a **single endpoint**: `POST /`
- The request must include a JSON body with the following fields:
  - `Type`: one of `"length"`, `"weight"`, or `"temperature"`
  - `From`: the unit to convert from
  - `To`: the unit to convert to
  - `Value`: the numeric value to convert
- The server performs the conversion and returns a message with the result

---

## ✅ Example Request

```http
POST / HTTP/1.1
Content-Type: application/json

{
  "Type": "length",
  "From": "meter",
  "To": "kilometer",
  "Value": 2000
}
```

## 🚀 Getting Started
1. Clone the repository

2. Run the server:
```bash
go run main.go
```
3. Send a POST request to http://localhost:8080/ with the required JSON body