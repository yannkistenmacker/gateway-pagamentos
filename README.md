# Gateway de Pagamento - Imersão Full Cycle

## Sobre o Projeto

Este projeto foi desenvolvido durante a [Imersão Full Stack & Full Cycle](https://imersao.fullcycle.com.br/evento/), onde construímos um Gateway de Pagamento completo utilizando arquitetura de microsserviços.

O objetivo é demonstrar a construção de um sistema distribuído moderno, com separação de responsabilidades, comunicação assíncrona e análise de fraudes em tempo real.

## Arquitetura

[Visualize a arquitetura completa aqui](https://link.excalidraw.com/readonly/Nrz6WjyTrn7IY8ZkrZHy)

### Componentes do Sistema

- **Frontend (Next.js)**
  - Interface do usuário para gerenciamento de contas e processamento de pagamentos
  - Desenvolvido com Next.js para garantir performance e boa experiência do usuário

- **Gateway (Go)**
  - Sistema principal de processamento de pagamentos
  - Gerencia contas, transações e coordena o fluxo de pagamentos
  - Publica eventos de transação no Kafka para análise de fraude

- **Apache Kafka**
  - Responsável pela comunicação assíncrona entre API Gateway e Antifraude
  - Garante confiabilidade na troca de mensagens entre os serviços
  - Tópicos específicos para transações e resultados de análise

- **Antifraude (Nest.js)**
  - Consome eventos de transação do Kafka
  - Realiza análise em tempo real para identificar possíveis fraudes
  - Publica resultados da análise de volta no Kafka

## Fluxo de Comunicação

1. Frontend realiza requisições para a API Gateway via REST
2. Gateway processa as requisições e publica eventos de transação no Kafka
3. Serviço Antifraude consome os eventos e realiza análise em tempo real
4. Resultados das análises são publicados de volta no Kafka
5. Gateway consome os resultados e finaliza o processamento da transação 


# Payment Gateway - English

## About the Project

This project was developed during the [Full Stack & Full Cycle Immersion](https://imersao.fullcycle.com.br/evento/), where we built a complete **Payment Gateway** using a **microservices architecture**.

The goal is to demonstrate the construction of a **modern distributed system**, with **separation of concerns**, **asynchronous communication**, and **real-time fraud analysis**.

## Architecture

[View the complete architecture here](https://link.excalidraw.com/readonly/Nrz6WjyTrn7IY8ZkrZHy)

### System Components

- **Frontend (Next.js)**
  - User interface for account management and payment processing  
  - Built with Next.js to ensure high performance and a great user experience

- **Gateway (Go)**
  - Main system responsible for payment processing  
  - Manages accounts, transactions, and coordinates the payment flow  
  - Publishes transaction events to Kafka for fraud analysis

- **Apache Kafka**
  - Responsible for asynchronous communication between the API Gateway and the Fraud Detection service  
  - Ensures reliable message exchange between services  
  - Specific topics for transactions and analysis results

- **Fraud Detection (Nest.js)**
  - Consumes transaction events from Kafka  
  - Performs real-time analysis to detect potential fraud  
  - Publishes the analysis results back to Kafka

## Communication Flow

1. The frontend sends requests to the API Gateway via REST  
2. The Gateway processes the requests and publishes transaction events to Kafka  
3. The Fraud Detection service consumes the events and performs real-time analysis  
4. The analysis results are published back to Kafka  
5. The Gateway consumes the results and completes the transaction processing
