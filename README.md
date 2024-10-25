# 🏥 Prescription Analysis API

An intelligent API endpoint that analyzes doctor's prescriptions using Google's Gemini LLM to extract medication details, identify potential causes, and assess fatality risks.

## 🚀 Features

- Extract medication names and dosage information from prescription images
- Identify potential causes of the illness
- Assess fatality risk factor
- Returns structured, easy-to-use JSON response

## 📝 API Specification

### Endpoint

```
POST /api/v1/drugs
```

### Request

- Method: `POST`
- Content-Type: `multipart/form-data`
- Body: 
  - `prescription` (file): Image file of the prescription

```bash
curl -X POST -F "file=@/path/to/file.png" http://localhost:8080/api/v1/drugs
```

### Response

```typescript
{
  medicine: {
    name: string;
    dosage: string;
  }[];
  causes: (
    "infectious" |
    "bad food" |
    "weather change" |
    "polluted air" |
    "polluted water" |
    "vectors"
  )[];
  fatality: number;
}
```

#### Example Response

```json
{
  "medicine": [
    {
      "name": "Amoxicillin",
      "dosage": "500mg twice daily"
    },
    {
      "name": "Paracetamol",
      "dosage": "650mg as needed"
    }
  ],
  "causes": ["infectious", "weather change"],
  "fatality": 0.2
}
```

## 🛠️ Technology Stack

- Backend: Golang (with CHI as the router)
- LLM: Google Gemini

## 🚦 Getting Started

1. Clone the repository
```bash
git clone https://github.com/VinukaThejana/getdrugs
```

2. Install dependencies
```bash
go mod tidy
```

3. Set up environment variables
```bash
GEMINI_API_KEY=your-api-key
PORT=your-port
ENV={oneof [development, production]}
```

4. Start the server
```bash
go run cmd/main.go
```


## 🤝 Contributing

Contributions, issues, and feature requests are welcome! Feel free to check [issues page](https://github.com/VinukaThejana/getdrugs/issues).

## ⚡ Performance

- Average response time: [to be determined]
- Accuracy rate: [to be determined]

## ⚠️ Limitations

- Currently supports prescriptions in English only
- Image must be in proper orientation, well lit and clear for optimal results
- Results should be verified by healthcare professionals
