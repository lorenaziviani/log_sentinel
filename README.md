# Log Sentinel

Monitoramento inteligente de logs com detecção de anomalias baseada em machine learning.

---

## Visão Geral

O Log Sentinel é uma solução para ingestão, armazenamento e análise de logs, com detecção de anomalias baseada em machine learning. O pipeline é composto por:

- **LogCollector (Go):** recebe logs via HTTP e arquivos locais, envia para ElasticSearch ou salva localmente.
- **ElasticSearch:** armazenamento e consulta eficiente dos logs.
- **AnomalyDetector (Python):** detecta anomalias nos logs usando ML.

---

## Arquitetura

```mermaid
flowchart LR
    A["App"] --> B["LogCollector (Go)"]
    F["Log File"] --> B
    B --> C["ElasticSearch"]
    B -.-> D["Local Fallback"]
    C --> E["AnomalyDetector (Python)"]
```

**Componentes:**

- **App:** Origem dos logs (serviços, aplicações)
- **Log File:** Arquivos locais de log
- **LogCollector (Go):** Coleta logs via HTTP e arquivos locais, envia para ElasticSearch ou salva localmente
- **ElasticSearch:** Armazenamento e consulta
- **Local Fallback:** Armazenamento local caso ElasticSearch esteja indisponível
- **AnomalyDetector (Python):** Detecção de anomalias

---

## Tech Stack

- Go (coletor e pipeline de ingestão)
- Python (scikit-learn para ML)
- gRPC ou REST (comunicação entre Go e Python)
- ElasticSearch (armazenamento e consulta de logs)

---

## Estrutura do Repositório

- `cmd/collector` — Serviço Go para coleta de logs
- `cmd/ml` — Serviço Python para detecção de anomalias
- `pkg/anomaly` — Lógica de detecção de anomalias em Go
- `internal/parser` — Parsing e normalização de logs
- `infra/elastic` — Scripts/configuração do ElasticSearch
- `docs/` — Documentação e diagramas

---

## Configuração

1. Copie o arquivo `.env.sample` para `.env` e ajuste as variáveis conforme necessário:

```env
# Endereço do ElasticSearch (padrão: http://localhost:9200)
ELASTIC_ADDR=http://localhost:9200

# Nome do índice no ElasticSearch (padrão: logs-sentinel)
ELASTIC_INDEX=logs-sentinel

# Diretório monitorado para arquivos locais de log (padrão: /var/log/log_sentinel)
LOG_SENTINEL_DIR=/var/log/log_sentinel
```

---

## Como Executar o LogCollector

```sh
cd cmd/collector
# Exporte as variáveis de ambiente ou use um gerenciador de env (ex: direnv, dotenv)
go run main.go
```

O serviço ficará disponível em `http://localhost:8080/logs` para receber logs via POST.

---

## Exemplo de Envio de Log via HTTP

```sh
curl -X POST http://localhost:8080/logs \
  -H 'Content-Type: application/json' \
  -d '{
    "timestamp": "2024-06-01T12:00:00Z",
    "level": "INFO",
    "message": "User login successful",
    "source": "auth-service"
  }'
```

---

## Observações

- O coletor faz parsing dos campos: `timestamp`, `level`, `message`, `source`.
- Logs podem ser enviados por múltiplas fontes (HTTP ou arquivos locais `.jsonl`).
- Se o ElasticSearch estiver indisponível, os logs são salvos localmente em `logs.jsonl`.
- O diretório monitorado para arquivos locais pode ser configurado via variável de ambiente.

---
