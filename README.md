# eCommerce Management System API

#### eFishery Academy Aqua Developer Batch 2 - Prasetya Ikra Priyadi

This repository contains a Go Language project as a part of Final project eFishery Academy Aqua Developer Batch 2, October 2022, Bandung, Indonesia
An API with eCommerce fungsionality including user management, product management, and transaction process (cart and order).
Thank you for visiting.

Build it 100% Go Language Programming. Still need improvement in terms of clean code and effectivity.

## Request For Comments (RFC)

You can access the RFC file on [RFC Documentation][rfc]

## Entity Relationship Diagram ERD

See how the database relationship works for this API [HERE][erd]

## Features

- Managing User with Register, Login, and Profiles
- Managing Product with Catalog
- Managing Transaction process with cart and order

## Motivation and Problems

An eCommerce application requires a backend service in the form of RestAPI to help manage data consumption in the client application. The services that are built are expected to accommodate the following functions:

#### Admin Privileges

Doing full access to the management of the product list that will be displayed on the eCommerce page includes:

- Add and manage detailed information about products
- Manage stock of added products
- Remove a product from the list of products that have been added

#### Public Access

Displays a list of products that have been added and can be viewed by the public. Public users can see the list and details of the product but do not have access to carry out transactions before logging in

#### User / Customer Privileges

Users who have logged in and have been authenticated can access as authorized users such as Profile Dashboard, Shopping Cart, Order, and Payment.

> The Backend service that is built at this stage is a basic service that may not be able to fully accommodate all eCommerce functions.

## Prerequisite

To run this app, you might to ensure youe machine has these instance installed:

1. Go Programming Language (Latest)
2. Docker
3. Postman or other RestAPI client
4. Code / Text Editor

## Installation

#### Clone repository

Clone the latest version of my repository

```sh
git clone https://github.com/prasetyaikrap/ems-aquadev.git
```

#### Modify docker-compose

There will not be too many changes that will be applied. But make sure docker compose is ready. Modify the value of the environment based on your preferences.

```sh
version: "3.9"
services:
  ems-postgres:
    container_name: ems-postgres
    image: prasetyaip/ems-postgresql:v1
    restart: always
    ports:
      - "5432:5432"
    environment:
      - POSTGRES_USER=USER ##Changes with your preferred name
      - POSTGRES_PASSWORD=PASSWORD ##Changes with your preferred password
      - POSTGRES_DB=DBNAME ##Changes with your preferred database name
    networks:
      - ems_network
    healthcheck:
      test: ["CMD-SHELL", "pg_isready -U postgres"]
      interval: 10s
      timeout: 5s
      retries: 5

  ems-apiserver:
    container_name: ems-apiserver
    image: prasetyaip/ems-aquadev:v3
    ports:
      - "1323:1323"
    environment:
      - DB_HOST=HOST ##Change HOST to database container name (in this case 'ems-postgres')
      - DB_USER=USER ## Same as Database User
      - DB_PASSWORD=PASSWORD ## Same as Database Password
      - DB_NAME=DBNAME ## Same as Database Name
      - DB_PORT=PORT ## Port for application (1323)
      - DB_TIMEZONE=TIMEZONE ##your preferences (ex: Asia/Jakarta)
      - JWT_SECRET=SECRET ##Scret key required for JWT
    networks:
      - ems_network
    depends_on:
      ems-postgres:
        condition: service_healthy
networks:
  ems_network:
    name: ems_network
```

#### Run Docker Compose

Execute docker compose in terminal with

```sh
docker-compose up
```

#### Check the connection with Postman

check the connection with API. In addtion, you can download the Postman API collection [Here][pcl]

## API Documentation

You can see the API Documentation [HERE][pdo]

## Thank You for visiting

[//]: #
[rfc]: https://docs.google.com/document/d/1g3BCPQUah1AR4N2S51iO1y7eJnkww0Pf7B5RcCmjh28/edit?usp=sharing
[pcl]: https://drive.google.com/file/d/1nE18FUeMWkS49yxjcCO3SpmcyRgfB1lK/view?usp=sharing
[pdo]: https://documenter.getpostman.com/view/21607478/2s83zdwRoW
[erd]: https://dbdiagram.io/d/633980097b3d2034ff063e5b
