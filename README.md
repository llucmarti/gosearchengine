# Search Engine

## Description
This is a simple search engine built for educational purposes.

## Prerequisites

Before you begin, ensure you have met the following requirements:

* You have installed the latest version of docker.

## Technologies Used

This project uses the following technologies:

* `Golang`: In this project, Go is used to handle HTTP requests, manage database connections, and implement the business logic.
* `PostgresSQL`: In this project, PostgreSQL is used to store and retrieve data efficiently. It works well with Go, making it a good choice for the database.

## API Endpoints

Here are the API endpoints available in this project:

* `GET /api/ads`: Searches for `{term}` materials and returns matching results.
  * Parameters:
    * `term`: The search term to look for.
    * `perPage`: Elements per page , Integer
    * `nPage`: Page to show, Integer`

* `GET /api/details`: Search for `{id}` ad, return it and its related material ads 
  * Parameters:
    * `id`: Advertising ID

## Data Structure

The data in this project is organized into three tables: `Product`, `Material`, and `Product_Material`.

* `Product`: This table represents the products in our system. Each product has an ID, a name, amount and price.

* `Material`: This table represents the materials that can be associated with a product. Each material has an ID and name

* `Product_Material`: This table represents the many-to-many relationship between products and materials. Each record in this table links a product with a material. This allows us to associate multiple materials with a product and vice versa.

This structure allows us to efficiently query our data and perform operations like finding all materials associated with a product, or finding all products that use a certain material.

## Contact

If you want to contact me you can reach me at `lluc.marti.calveres@gmail.com`.