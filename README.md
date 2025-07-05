# 🔍 Log Sentinel - Observabilidade & Anomalias em Logs

<div align="center">
<img src=".gitassets/cover.png" width="350" />

<div data-badges>
  <img src="https://img.shields.io/github/stars/lorenaziviani/log_sentinel?style=for-the-badge&logo=github" alt="GitHub stars" />
  <img src="https://img.shields.io/github/forks/lorenaziviani/log_sentinel?style=for-the-badge&logo=github" alt="GitHub forks" />
  <img src="https://img.shields.io/github/last-commit/lorenaziviani/log_sentinel?style=for-the-badge&logo=github" alt="GitHub last commit" />
</div>

<div data-badges>
  <img src="https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white" alt="Go" />
  <img src="https://img.shields.io/badge/Python-3776AB?style=for-the-badge&logo=python&logoColor=white" alt="Python" />
  <img src="https://img.shields.io/badge/ElasticSearch-005571?style=for-the-badge&logo=elasticsearch&logoColor=white" alt="ElasticSearch" />
  <img src="https://img.shields.io/badge/Prometheus-E6522C?style=for-the-badge&logo=prometheus&logoColor=white" alt="Prometheus" />
  <img src="https://img.shields.io/badge/Grafana-F46800?style=for-the-badge&logo=grafana&logoColor=white" alt="Grafana" />
  <img src="https://img.shields.io/badge/Docker-2496ED?style=for-the-badge&logo=docker&logoColor=white" alt="Docker" />
  <img src="https://img.shields.io/badge/FastAPI-009688?style=for-the-badge&logo=fastapi&logoColor=white" alt="FastAPI" />
</div>
</div>

O **Log Sentinel** é uma plataforma de observabilidade e detecção de anomalias em logs, combinando ingestão de logs, machine learning, ElasticSearch, Prometheus e Grafana para monitoramento inteligente e reação automática a incidentes.

✔️ **Ingestão de logs multi-fonte** (HTTP, arquivos)

✔️ **Detecção de anomalias** em tempo real via ML (Isolation Forest)

✔️ **Alertas automáticos** em caso de picos de anomalia (Discord)

✔️ **Observabilidade completa** com Prometheus e Grafana

✔️ **Dashboards ricos** para logs, anomalias, métricas e latência

✔️ **Fallback local** e rastreabilidade ponta-a-ponta

Desenvolvido em Go e Python, pronto para produção, extensível e fácil de integrar.

---

## 🖥️ Como rodar este projeto

### Requisitos:

- [Go 1.24+](https://golang.org/doc/install)
- [Python 3.10+](https://www.python.org/)
- [Docker & Docker Compose](https://docs.docker.com/get-docker/)

### Execução:

1. Clone este repositório:
   ```sh
   git clone https://github.com/lorenaziviani/log_sentinel.git
   cd log_sentinel
   ```
2. Instale dependências Go e Python:
   ```sh
   go mod download
   cd cmd/ml && pip install -r requirements.txt
   ```
3. Configure variáveis de ambiente:
   ```sh
   cp .env.sample .env
   # Edite .env conforme necessário
   ```
4. Suba todos os serviços com Docker Compose:
   ```sh
   docker-compose up --build
   ```
5. Ou execute localmente:
   ```sh
   # ElasticSearch e Kibana via Docker
   docker-compose up -d elasticsearch kibana
   # ML
   cd cmd/ml && uvicorn main:app --host 0.0.0.0 --port 8000
   # Coletor
   cd ../collector && go run main.go
   ```
6. Envie logs de teste:
   ```sh
   make integration-test
   # ou
   bash cmd/collector/logs-test.sh
   ```
7. Acesse os serviços:
   - **Kibana**: [http://localhost:5601](http://localhost:5601)
   - **Grafana**: [http://localhost:3000](http://localhost:3000)
   - **Prometheus**: [http://localhost:9090](http://localhost:9090)
   - **ML API**: [http://localhost:8000/docs](http://localhost:8000/docs)

---

## 📝 Features do projeto

🔎 **Ingestão & Parsing**

- Recebe logs via HTTP e arquivos locais
- Parsing centralizado e normalização de campos
- Suporte a múltiplas fontes e formatos

🤖 **Detecção de Anomalias (ML)**

- Serviço Python com Isolation Forest
- Treinamento e predição via API REST
- Score de anomalia e classificação em tempo real

🚨 **Alertas & Reação**

- Geração automática de alertas (`level: ALERT`) em caso de picos de anomalia
- Persistência de alertas e logs anômalos em índices dedicados
- Integração com Discord via webhook

📊 **Observabilidade Completa**

- Métricas Prometheus: volume, anomalias, latência do ML
- Dashboards Grafana: volume, % anomalia, tempo de resposta, alertas
- Dashboards Kibana: logs, anomalias, rastreabilidade

🛠️ **Administração & Testes**

- Makefile com targets para testes, lint, execução e integração
- Script de integração real (`logs-test.sh`)
- Fallback local automático se ElasticSearch indisponível

---

## 🛠️ Comandos de Teste

```bash
# Testes unitários do collector
make test

# Lint
make lint

# Executar o coletor
make run

# Teste de integração (envia logs reais)
make integration-test

# Treinar modelo ML
make train-ml

# Testar endpoint de predição ML
make predict-ml

# Testes unitários do ml
make test-ml
```

---

## 📈 Monitoramento e Dashboards

### Grafana Dashboard

Acesse [http://localhost:3000](http://localhost:3000) para visualizar:

- Volume de logs
- % de anomalias
- Latência do ML
- Alertas em tempo real

![Dashboard Grafana](.gitassets/grafana.png)

### Kibana

Acesse [http://localhost:5601](http://localhost:5601) para:

- Explorar logs e anomalias
- Criar dashboards customizados
- Rastrear logs do recebimento à classificação

![Dashboard Elastic](.gitassets/elastic-logs.png)
![Dashboard Elastic](.gitassets/elastic-dash.png)

### Prometheus Metrics

Acesse [http://localhost:9090](http://localhost:9090) para monitorar:

- Métricas em tempo real do Log Sentinel (coletor, ML, anomalias)
- Targets e endpoints monitorados (serviços Go, ML, etc)
- Queries customizadas para análise de volume de logs, latência do ML, % de anomalias
- Alertas e regras configuradas para detecção de picos, falhas ou anomalias

![Prometheus UI - Targets](.gitassets/prometheus.png)

### Alertas

- Envio de alertas de anomalias no Discord

![Prometheus UI - Targets](.gitassets/discord.png)

---

## 🏗️ Arquitetura do Sistema

![Architecture](docs/architecture.drawio.png)

**Fluxo detalhado:**

1. Recebe log (HTTP ou arquivo)
2. Salva no ElasticSearch (ou local)
3. Consulta ML para detecção de anomalia
4. Salva anomalias e alertas em índices dedicados
5. Exposição de métricas Prometheus
6. Dashboards em Grafana e Kibana
7. Alertas enviados para Discord

---

## 🌐 Variáveis de Ambiente

```env
# .env.example
ELASTIC_ADDR=http://elasticsearch:9200
ML_URL=http://ml:8000/predict
DISCORD_WEBHOOK_URL=https://discord.com/api/webhooks/...
```

---

## 📁 Estrutura de Pastas

```
log_sentinel/
  go.mod
  go.sum
  docker-compose.yml
  Makefile
  .env.sample
  cmd/
    collector/      # Projeto Go (main.go, notifier.go, logs-test.sh, etc)
    ml/             # Projeto Python (main.py, requirements.txt, etc)
  internal/
    parser/         # Pacotes Go internos
  pkg/
    anomaly/        # Outros pacotes Go
  docs/
    architecture.drawio.png
  .gitassets/       # Imagens para README
```

---

## 💎 Links úteis

- [Go Documentation](https://golang.org/doc/)
- [FastAPI](https://fastapi.tiangolo.com/)
- [ElasticSearch](https://www.elastic.co/)
- [Prometheus](https://prometheus.io/)
- [Grafana](https://grafana.com/)
- [Docker](https://www.docker.com/)

---
