# nemmy

Government of the nemmy, by the nemmy, for the nemmy.

## URL

<https://nemmy.vercel.app>

## Project Structure

| Role                   | Link                                     |
| ---------------------- | ---------------------------------------- |
| Front-end              | <https://nemmy.vercel.app>               |
| Back-end               | <https://github.com/ynm3n/nemmy-backend> |
| API Documentation      | <https://nemmy-swagger.pages.dev>        |
| Database Documentation | <https://dbdocs.io/light-planck/nemmy>   |

## Setup

> [!NOTE]
> We use [mise](https://mise.jdx.dev/) to manage the Node.js version.
> Also, we fix the [pnpm](https://pnpm.io/) version using [corepack](https://github.com/nodejs/corepack), so please install the compatible version.

```zsh
mise install
corepack enable
corepack use pnpm
pnpm i
```

### Starting the Local Server

#### Front-end

```zsh
pnpm dev:front
```

#### Swagger

```zsh
pnpm dev:swagger
```
