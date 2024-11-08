# law-finder

## Usage

<details>
<summary><code>GET</code> <code><b>/</b></code></summary>

##### Responses

> | http code | content-type | response      |
> | --------- | ------------ | ------------- |
> | `200`     | `text/plain` | `Hello World` |

</details>

<details>
<summary><code>GET</code> <code><b>/law-finder</b></code> <code>(Get raw data of the specific act)</code></summary>

##### Query Parameters

> | key | required | data type | description           |
> | --- | -------- | --------- | --------------------- |
> | law | true     | string    | Article to search for |

##### Responses

> | http code    | content-type | response           |
> | ------------ | ------------ | ------------------ |
> | `200`        | `text/plain` | `Act's plain text` |
> | `400`, `500` | `text/plain` | `error message`    |

</details>
<details>
<summary><code>POST</code> <code><b>/law-finder</b></code> <code>(Get raw data of specifc act's specific article)</code></summary>

##### Query Parameters

> | key | required | data type | description           |
> | --- | -------- | --------- | --------------------- |
> | law | true     | string    | Article to search for |

##### Body (application/json)

> | key     | required | data type | description                                          |
> | ------- | -------- | --------- | ---------------------------------------------------- |
> | article | true     | string    | Can be formatted as `第一條`, `第 1 條`, `一` or `1` |

##### Responses

> | http code | content-type | response               |
> | --------- | ------------ | ---------------------- |
> | `200`     | `text/plain` | `Article's plain text` |
> | `400`     | `text/plain` | `error message`        |

</details>

## Example

1.  Retrieving all 道路交通安全規則

    ```curl
    curl -X 'GET' http://localhost:8080?law=道路交通安全規則
    ```

2.  Retrieving specifc article from 道路交通安全規則

    ```curl
    curl -X 'POST' 'http://localhost:8080/law-finder?law=道路交通安全規則' -H 'Content-Type:application/json' -d '{"article": "第十五條"}'
    ```
