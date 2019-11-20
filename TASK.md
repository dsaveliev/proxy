Given
-----

You are given an external API endpoint which allows to query recipe information. Each recipe can be accessed by an `integer` id.
The recipe id enumeration starts from `1`.

Example HTTP calls

```
curl -X GET https://example.com/endpoint/1
curl -X GET https://example.com/endpoint/2
curl -X GET https://example.com/endpoint/5
```

Task
----

Design an application which would act as a reverse proxy and expose the _aggregated recipes_ from the external API over HTTP.

#### Requirements

- The recipes in the aggregated list **must** contain the same data as the original recipes. Data modifications are **not allowed**.
- The endpoint response **must** be `JSON` encoded.
- The endpoint response time **must** be lower than `1s`.
- The application should be stateless, i.e. it is **not allowed** to cache the recipe response on the application side.
- The endpoint **should not** render all the recipes in a single response. It is **allowed** to make [use of pagination](http://docs.oasis-open.org/odata/odata/v4.01/cs01/part2-url-conventions/odata-v4.01-cs01-part2-url-conventions.html#_Toc505773300).

##### Use Case #1 - all recipes

A user should be able to retrieve an aggregated list of **all the recipes** from the source API.

_Specific requirements_
- The endpoint **must** provide access to ALL available recipes.
- The order for the rendered recipes **is irrelevant**
- The solution should operate under the assumption that the source API contains an unlimited number of recipes.

> `All available recipes` are the recipes with the `id` lower than the `id` with the first `404 Not Found` HTTP response status code.
>
> For example, if
>  `curl -X GET https://example.com/endpoint/99999` returns HTTP status code `200 OK`
> and
> `curl -X GET https://example.com/endpoint/100000` returns HTTP status code `404 Not Found`
> then
> `all available recipes` are the ones with the `ids` from 1 to 99999



Example endoint: `GET http://myservice.io/recipes`

```json
[
    {
        "id": "5",
        // ...
    },
    {
        "id": "1",
        // ...
    },
    {
        "id": "2",
        // ...
    }
]
```

##### Use Case #2 - recipes by `id`

A user should be able to retrieve a list of **aggregated recipes** from the source API by a given `id`.

_Specific requirements_

- The endpoint **must** provide access to the recipes by the provided `id`.
- The recipes should be ordered by `prepTime` from lowest to highest.

Example endpoint and response: `GET http://myservice.io/recipes?ids=1,2,5`

```json
[
    {
        "id": "1",
        "prepTime": "PT30",
        // ...
    },
    {
        "id": "5",
        "prepTime": "PT30",
        // ...
    },
    {
        "id": "2",
        "prepTime": "PT35",
        // ...
    }
]
```

Evaluation Criteria
--------------

1. The problems are solved efficiently and effectively, the application works as expected.
2. The application is supplied with the setup scripts. Consider using docker and a one-liner setup.
3. You demonstrate the knowledge on how to test the critical parts of the application. We **do not require** 100% coverage.
4. The application is well and logically organised.
5. The submission is accompanied by documentation with the reasoning on the decisions taken.
6. The code is documented and is easy to follow.
7. The answers you provide during code review.
8. An informative, detailed description in the PR.
9. Following the industry standard style guide.
10. A git history (even if brief) with clear, concise commit messages.

---

Good luck!
