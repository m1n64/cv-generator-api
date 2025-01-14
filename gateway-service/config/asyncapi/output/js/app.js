
    const schema = {
  "asyncapi": "3.0.0",
  "info": {
    "title": "Simple CV Generator WebSocket API",
    "version": "1.0.0",
    "description": "This WebSocket API sends notifications and updates to clients about CV generation tasks. \nIt requires a `token` parameter for authentication to connect.\n",
    "externalDocs": {
      "url": "https://api.resumego.online/docs/",
      "description": "Full API documentation"
    }
  },
  "servers": {
    "local-private": {
      "host": "localhost:8000",
      "pathname": "/ws/private?token={token}",
      "protocol": "ws",
      "description": "Local development server for testing WebSocket connections.",
      "variables": {
        "token": {
          "description": "Token for authentication.",
          "default": "your-token"
        }
      }
    },
    "production-private": {
      "host": "api.resumego.online",
      "pathname": "/ws/private?token={token}",
      "protocol": "wss",
      "description": "Production WebSocket server for private notifications.",
      "variables": {
        "token": {
          "description": "Token for authentication.",
          "default": "your-token"
        }
      }
    }
  },
  "channels": {
    "private": {
      "address": "private",
      "messages": {
        "subscribe.message": {
          "summary": "Server-to-client message.",
          "description": "This message is sent from the server to the client, typically to notify about \nthe status of a task or other important updates.\n",
          "contentType": "application/json",
          "payload": {
            "type": "object",
            "required": [
              "type",
              "user_id",
              "message"
            ],
            "properties": {
              "type": {
                "type": "string",
                "description": "The type of message (e.g., success, error).",
                "example": "success",
                "x-parser-schema-id": "<anonymous-schema-2>"
              },
              "user_id": {
                "type": "string",
                "format": "uuid",
                "description": "Unique identifier of the user receiving the message.",
                "example": "2f89414f-9e64-434c-908a-66b529d558b9",
                "x-parser-schema-id": "<anonymous-schema-3>"
              },
              "message": {
                "type": "string",
                "description": "A descriptive message about the update or event.",
                "example": "pdf.generated.success",
                "x-parser-schema-id": "<anonymous-schema-4>"
              }
            },
            "x-parser-schema-id": "<anonymous-schema-1>"
          },
          "x-parser-unique-object-id": "subscribe.message",
          "x-parser-message-name": "serverMessage"
        }
      },
      "description": "Channel for receiving private messages and updates.",
      "x-parser-unique-object-id": "private"
    }
  },
  "operations": {
    "private.subscribe": {
      "action": "send",
      "channel": "$ref:$.channels.private",
      "summary": "Listen to server messages.",
      "description": "The client subscribes to this channel to receive server-sent notifications, \nsuch as success messages and task updates.\n",
      "messages": [
        "$ref:$.channels.private.messages.subscribe.message"
      ],
      "x-parser-unique-object-id": "private.subscribe"
    }
  },
  "components": {
    "messages": {
      "serverMessage": "$ref:$.channels.private.messages.subscribe.message"
    },
    "schemas": {
      "serverMessagePayload": {
        "type": "object",
        "properties": {
          "type": {
            "type": "string",
            "description": "Message type (e.g., success, error).",
            "example": "success",
            "x-parser-schema-id": "<anonymous-schema-5>"
          },
          "user_id": {
            "type": "string",
            "description": "UUID of the user.",
            "example": "2f89414f-9e64-434c-908a-66b529d558b9",
            "x-parser-schema-id": "<anonymous-schema-6>"
          },
          "message": {
            "type": "string",
            "description": "Detailed message about the server event.",
            "example": "pdf.generated.success",
            "x-parser-schema-id": "<anonymous-schema-7>"
          }
        },
        "x-parser-schema-id": "serverMessagePayload"
      }
    }
  },
  "x-parser-spec-parsed": true,
  "x-parser-api-version": 3,
  "x-parser-spec-stringified": true
};
    const config = {"show":{"sidebar":true},"sidebar":{"showOperations":"byDefault"}};
    const appRoot = document.getElementById('root');
    AsyncApiStandalone.render(
        { schema, config, }, appRoot
    );
  