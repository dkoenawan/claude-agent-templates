# Domain Specialization Guide

This guide explains how to create and work with domain-specific agents in Claude Agent Templates, covering technology-specific patterns, frameworks, and best practices.

## Table of Contents

- [Overview](#overview)
- [Domain-Specific Patterns](#domain-specific-patterns)
- [Technology Stack Guidelines](#technology-stack-guidelines)
- [Architecture Patterns](#architecture-patterns)
- [Creating New Domain Agents](#creating-new-domain-agents)
- [Cross-Domain Considerations](#cross-domain-considerations)

## Overview

Domain specialization allows agents to provide technology-specific expertise while maintaining consistency across the development workflow. Each domain has specialized agents for architecture, implementation, and testing that understand the nuances of their respective technology stacks.

### Available Domains

- **Core** - Language-agnostic workflow coordination
- **Python** - Python ecosystem (FastAPI, Django, Flask)
- **.NET** - Microsoft .NET ecosystem (ASP.NET Core, Entity Framework)
- **Node.js** - JavaScript/TypeScript ecosystem (Express.js, React, TypeScript)
- **Java** - Java ecosystem (Spring Boot, Spring Framework, JPA)

### Agent Roles per Domain

Each domain (except core) has three specialized agents:
1. **Solution Architect** - Architecture planning and design
2. **Software Engineer** - Implementation and development
3. **Test Engineer** - Testing strategy and implementation

## Domain-Specific Patterns

### Python Domain

**Frameworks and Tools:**
- **Web Frameworks**: FastAPI (preferred), Django, Flask
- **Data Layer**: SQLAlchemy, Alembic, Pydantic
- **Testing**: pytest, unittest, coverage, mock
- **Quality**: black, ruff, mypy, pre-commit
- **Async**: asyncio, aiohttp, async/await patterns

**Architecture Patterns:**
- **Hexagonal Architecture** with clear port/adapter separation
- **Domain-Driven Design** with bounded contexts
- **Dependency Injection** using dependency-injector or similar
- **Event-Driven Architecture** with async messaging

**File Structure Example:**
```
src/
├── domain/
│   ├── entities/
│   ├── value_objects/
│   └── services/
├── application/
│   ├── use_cases/
│   ├── interfaces/
│   └── services/
├── infrastructure/
│   ├── repositories/
│   ├── adapters/
│   └── external/
└── presentation/
    ├── api/
    ├── schemas/
    └── middleware/
```

**Best Practices:**
- Use type hints throughout codebase
- Follow PEP 8 and PEP 257 standards
- Implement comprehensive async patterns
- Use Pydantic for data validation
- Maintain >80% test coverage

### .NET Domain

**Frameworks and Tools:**
- **Web Frameworks**: ASP.NET Core, Blazor, Web API
- **Data Layer**: Entity Framework Core, Dapper
- **Testing**: xUnit, NUnit, MSTest, Moq, FluentAssertions
- **Quality**: SonarAnalyzer, StyleCop, EditorConfig
- **Build**: MSBuild, dotnet CLI

**Architecture Patterns:**
- **Clean Architecture** with strict layer separation
- **CQRS** with MediatR for complex domains
- **Repository Pattern** with Unit of Work
- **Dependency Injection** using built-in DI container

**Project Structure Example:**
```
Solution/
├── Domain/
│   ├── Entities/
│   ├── ValueObjects/
│   └── Services/
├── Application/
│   ├── UseCases/
│   ├── Interfaces/
│   └── Services/
├── Infrastructure/
│   ├── Persistence/
│   ├── External/
│   └── Configuration/
└── Presentation/
    ├── Controllers/
    ├── ViewModels/
    └── Middleware/
```

**Best Practices:**
- Follow SOLID principles rigorously
- Use nullable reference types
- Implement proper exception handling
- Use configuration patterns (IOptions)
- Leverage built-in logging and DI

### Node.js Domain

**Frameworks and Tools:**
- **Web Frameworks**: Express.js, Fastify, Koa
- **Type Safety**: TypeScript (strongly recommended)
- **Testing**: Jest, Mocha, Supertest, Cypress
- **Quality**: ESLint, Prettier, Husky
- **Build**: Webpack, Vite, esbuild

**Architecture Patterns:**
- **Hexagonal Architecture** with TypeScript interfaces
- **Event-Driven Architecture** with EventEmitter
- **Microservices** with proper service boundaries
- **Middleware Pattern** for cross-cutting concerns

**Project Structure Example:**
```
src/
├── domain/
│   ├── entities/
│   ├── valueObjects/
│   └── services/
├── application/
│   ├── useCases/
│   ├── ports/
│   └── services/
├── infrastructure/
│   ├── adapters/
│   ├── repositories/
│   └── external/
└── presentation/
    ├── controllers/
    ├── middleware/
    └── routes/
```

**Best Practices:**
- Use TypeScript for type safety
- Implement proper error handling
- Use async/await over promises
- Follow Node.js security best practices
- Implement proper logging and monitoring

### Java Domain

**Frameworks and Tools:**
- **Web Frameworks**: Spring Boot, Spring MVC, Spring WebFlux
- **Data Layer**: Spring Data JPA, Hibernate, MyBatis
- **Testing**: JUnit 5, Mockito, TestContainers, AssertJ
- **Quality**: SpotBugs, PMD, Checkstyle
- **Build**: Maven, Gradle

**Architecture Patterns:**
- **Clean Architecture** with Spring annotations
- **Hexagonal Architecture** with Spring Boot
- **Microservices** with Spring Cloud
- **Event Sourcing** with Spring Events

**Project Structure Example:**
```
src/main/java/
├── domain/
│   ├── model/
│   ├── service/
│   └── repository/
├── application/
│   ├── usecase/
│   ├── port/
│   └── service/
├── infrastructure/
│   ├── adapter/
│   ├── repository/
│   └── config/
└── presentation/
    ├── controller/
    ├── dto/
    └── config/
```

**Best Practices:**
- Use Spring Boot best practices
- Implement proper exception handling
- Use Java 17+ features appropriately
- Follow Spring Security patterns
- Implement comprehensive integration tests

## Technology Stack Guidelines

### Language-Specific Considerations

**Python:**
- Version: Python 3.11+ recommended
- Package Management: Poetry or pip with requirements.txt
- Virtual Environments: venv, conda, or Poetry
- Code Style: black, isort, flake8/ruff

**.NET:**
- Version: .NET 8+ LTS recommended
- Package Management: NuGet with PackageReference
- Project Templates: Use .NET templates
- Code Style: EditorConfig, StyleCop

**Node.js:**
- Version: Node.js 20+ LTS recommended
- Package Management: npm, yarn, or pnpm
- Type System: TypeScript strongly recommended
- Code Style: Prettier, ESLint

**Java:**
- Version: Java 17+ LTS recommended
- Build Tools: Maven or Gradle
- Framework: Spring Boot 3+
- Code Style: Google Java Style or similar

### Database Integration

**Python:**
- **PostgreSQL**: asyncpg, psycopg2, SQLAlchemy
- **MongoDB**: motor, pymongo
- **Redis**: redis-py, aioredis

**.NET:**
- **SQL Server**: Entity Framework Core
- **PostgreSQL**: Npgsql.EntityFrameworkCore.PostgreSQL
- **MongoDB**: MongoDB.Driver

**Node.js:**
- **PostgreSQL**: pg, Prisma, TypeORM
- **MongoDB**: mongoose, mongodb
- **Redis**: redis, ioredis

**Java:**
- **PostgreSQL**: Spring Data JPA, HikariCP
- **MongoDB**: Spring Data MongoDB
- **Redis**: Spring Data Redis

## Architecture Patterns

### Hexagonal Architecture Implementation

**Core Principles:**
1. **Domain Independence**: Business logic independent of external concerns
2. **Port and Adapter Pattern**: Clear interfaces for external dependencies
3. **Dependency Inversion**: Dependencies point inward toward domain
4. **Testability**: Easy to test with mock adapters

**Layer Responsibilities:**

**Domain Layer:**
- Entities and value objects
- Business rules and domain services
- Domain events
- No external dependencies

**Application Layer:**
- Use cases and application services
- Orchestration of domain operations
- Input/output ports (interfaces)
- Application-specific business rules

**Infrastructure Layer:**
- Database adapters
- External service clients
- Framework-specific implementations
- Technical concerns

**Presentation Layer:**
- HTTP controllers/handlers
- Request/response models
- Authentication/authorization
- Input validation

### Testing Strategy by Domain

**Unit Testing:**
- **Python**: pytest with fixtures and parametrized tests
- **.NET**: xUnit with FluentAssertions and AutoFixture
- **Node.js**: Jest with describe/it blocks and mocks
- **Java**: JUnit 5 with Mockito and AssertJ

**Integration Testing:**
- **Python**: TestContainers with pytest
- **.NET**: WebApplicationFactory with TestContainers
- **Node.js**: Supertest with test databases
- **Java**: Spring Boot Test with TestContainers

**End-to-End Testing:**
- **Python**: Playwright or Selenium
- **.NET**: Playwright or Selenium
- **Node.js**: Cypress or Playwright
- **Java**: Selenium or TestContainers

## Creating New Domain Agents

### Step 1: Define Domain Scope

1. Identify technology stack and frameworks
2. Define architectural patterns
3. Establish testing strategies
4. Document best practices

### Step 2: Create Agent Directory

```bash
mkdir -p agents/newdomain
```

### Step 3: Create Domain Agents

Create three agents following the naming convention:
- `solution-architect-newdomain.md`
- `software-engineer-newdomain.md`
- `test-engineer-newdomain.md`

### Step 4: Implement Agent Specifications

Use the standard specification format with domain-specific:
- Framework expertise
- Tool recommendations
- Architecture patterns
- Testing strategies
- Best practices

### Step 5: Add Domain Classification

Update `scripts/classify-domain.py` to include:
- Domain keywords
- File extensions
- Framework patterns
- Tool indicators

### Step 6: Create Tests

Add comprehensive tests in `tests/integration/`:
- `test_newdomain_agent_workflow.py`

### Step 7: Update Documentation

Update:
- README.md with new domain
- This guide with domain specifics
- Validation rules if needed

## Cross-Domain Considerations

### Shared Patterns

**Common Architecture Patterns:**
- Clean/Hexagonal Architecture
- Dependency Injection
- Repository Pattern
- Domain-Driven Design

**Common Quality Practices:**
- >80% test coverage
- Automated code formatting
- Static analysis
- Comprehensive documentation

### Integration Points

**Database Integration:**
- Consistent migration strategies
- Common ORM patterns
- Transaction management

**API Design:**
- RESTful principles
- OpenAPI documentation
- Consistent error handling
- Authentication/authorization

**Monitoring and Logging:**
- Structured logging
- Performance metrics
- Error tracking
- Health checks

### Technology Migration

When migrating between domains:
1. Assess current architecture patterns
2. Map equivalent frameworks and tools
3. Plan data migration strategy
4. Maintain API compatibility
5. Implement gradual migration

### Multi-Domain Projects

For projects spanning multiple domains:
1. Use core agents for coordination
2. Maintain consistent API contracts
3. Implement proper service boundaries
4. Use shared data formats
5. Coordinate deployment strategies