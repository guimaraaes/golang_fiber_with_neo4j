# golang_fiber_with_neo4j

> API para modelagem de entidades e relacionamentos para o Neo4j utilizando o framework web fiber baseado na linguagem golang

## Quick Start

Para iniciar a aplicação no seu computador é necessário ter o Docker instalado e seguir os seguintes passos:

```bash
# clone o repertório na sua máquina local
git clone github.com/guimaraaes/golang_fiber_with_neo4j.git

# Direcione-se para o diretório
cd golang_fiber_with_neo4j

# Construa a aplicação utilizando Docker
docker-compose build

# Compile a aplicação utilizando Docker
docker-compose up
```

## Descrição da modelagem

Nesta API a modelagem foi feita definindo variáveis do tipo struct no package model. Por exemplo, nesta aplicação há uma entidade Person, um relacionamento KNOW e um modelo genérico de como relacionar as entidades (Person-KNOWS->Person). Pelo diagrama é possível visualizar essa modelagem,

![img](https://raw.githubusercontent.com/guimaraaes/golang_fiber_with_neo4j/master/arrow-schema/arrows.svg)
