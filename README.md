# pro-profile-ai-api

## Load environment variables

1. Create `.env` file and populate it.
2. Execute command to load variables into shell

```bash
export $(grep -v '^#' .env | xargs)
```
