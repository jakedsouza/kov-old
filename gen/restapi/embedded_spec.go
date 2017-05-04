///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "# RESTful API for the Kubernetes on vSphere (KOV)\n",
    "title": "API Specification for the Kubernetes on vSphere (KOV)",
    "version": "0.1.0"
  },
  "paths": {
    "/clusters": {
      "get": {
        "description": "get a list of all clusters managed by the VCCS",
        "summary": "get a list of all clusters",
        "operationId": "listClusters",
        "parameters": [
          {
            "$ref": "#/parameters/requestId"
          }
        ],
        "responses": {
          "200": {
            "description": "200 response with the list of clusters",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/cluster"
              }
            }
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      },
      "post": {
        "description": "creates a cluster",
        "summary": "creates a cluster",
        "operationId": "createCluster",
        "parameters": [
          {
            "$ref": "#/parameters/requestId"
          },
          {
            "description": "the config of the cluster to be created",
            "name": "clusterConfig",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/clusterConfig"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "create cluster task has been accepted",
            "schema": {
              "$ref": "#/definitions/taskId"
            }
          },
          "409": {
            "$ref": "#/responses/errorClusterNameConflict"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      }
    },
    "/clusters/{name}": {
      "put": {
        "description": "updates a cluster with the given update config",
        "summary": "updates a cluster",
        "operationId": "updateCluster",
        "parameters": [
          {
            "$ref": "#/parameters/requestId"
          },
          {
            "type": "string",
            "x-nullable": false,
            "description": "the cluster name to be updated",
            "name": "name",
            "in": "path",
            "required": true
          },
          {
            "description": "the new config of the cluster to be updated",
            "name": "clusterUpdateConfig",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/clusterUpdateConfig"
            }
          }
        ],
        "responses": {
          "202": {
            "description": "update cluster task has been accepted",
            "schema": {
              "$ref": "#/definitions/taskId"
            }
          },
          "404": {
            "$ref": "#/responses/errorClusterNotFound"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      },
      "delete": {
        "description": "deletes a cluster with the given name",
        "summary": "deletes a cluster",
        "operationId": "deleteCluster",
        "parameters": [
          {
            "$ref": "#/parameters/requestId"
          },
          {
            "type": "string",
            "x-nullable": false,
            "description": "the cluster name to be deleted",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "202": {
            "description": "delete cluster task has been accepted",
            "schema": {
              "$ref": "#/definitions/taskId"
            }
          },
          "404": {
            "$ref": "#/responses/errorClusterNotFound"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      }
    },
    "/tasks/{taskid}": {
      "get": {
        "summary": "get the task for the given task id",
        "operationId": "getTask",
        "parameters": [
          {
            "$ref": "#/parameters/requestId"
          }
        ],
        "responses": {
          "200": {
            "description": "the task for the given task id",
            "schema": {
              "$ref": "#/definitions/task"
            }
          },
          "404": {
            "$ref": "#/responses/errorTaskNotFound"
          },
          "default": {
            "$ref": "#/responses/errorDefault"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "description": "the id for a task",
          "name": "taskid",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "cluster": {
      "type": "object",
      "required": [
        "config",
        "status"
      ],
      "properties": {
        "config": {
          "$ref": "#/definitions/clusterConfig"
        },
        "status": {
          "$ref": "#/definitions/clusterStatus"
        }
      }
    },
    "clusterConfig": {
      "type": "object",
      "required": [
        "name",
        "minNodes",
        "noOfMasters",
        "resourcePool",
        "managementNetwork"
      ],
      "properties": {
        "credentials": {
          "$ref": "#/definitions/credentials"
        },
        "managementNetwork": {
          "description": "the management network for the deployed nodes, will have ssh port enabled",
          "type": "string"
        },
        "masterSize": {
          "allOf": [
            {
              "$ref": "#/definitions/instanceSize"
            },
            {
              "default": "small"
            },
            {
              "description": "the size of the master nodes"
            }
          ]
        },
        "maxNodes": {
          "description": "the minimum number of nodes that can be deployed",
          "type": "integer",
          "format": "int32"
        },
        "minNodes": {
          "description": "the minimum number of nodes that can be deployed",
          "type": "integer",
          "format": "int32",
          "minimum": 1
        },
        "name": {
          "description": "the cluster name, should be valid for use in dns names",
          "type": "string",
          "maxLength": 63,
          "minLength": 3,
          "pattern": "^[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?(\\.[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?)*$"
        },
        "noOfMasters": {
          "description": "the number of master nodes to create",
          "type": "integer",
          "format": "int32",
          "default": 1,
          "minimum": 1
        },
        "nodeNetwork": {
          "description": "the network used for node-to-node communication, defaults to management network",
          "type": "string"
        },
        "nodeResourcePools": {
          "type": "array",
          "items": {
            "type": "string",
            "minLength": 1
          }
        },
        "nodeSize": {
          "allOf": [
            {
              "$ref": "#/definitions/instanceSize"
            },
            {
              "description": "the size of the worker nodes, when not   specified defaults to master size"
            }
          ]
        },
        "publicNetwork": {
          "description": "the public network to expose ports on, defaults to management network",
          "type": "string"
        },
        "resourcePool": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        },
        "storageClasses": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/storageClass"
          }
        },
        "thumbprint": {
          "description": "the thumbprint of the vCenter server certificate",
          "type": "string",
          "minLength": 57,
          "pattern": "[a-fA-F0-9:]+"
        }
      }
    },
    "clusterStatus": {
      "description": "the status of the cluster",
      "type": "string",
      "enum": [
        "active",
        "inactive",
        "error"
      ]
    },
    "clusterUpdateConfig": {
      "type": "object",
      "required": [
        "name",
        "minNodes",
        "noOfMasters"
      ],
      "properties": {
        "credentials": {
          "$ref": "#/definitions/credentials"
        },
        "maxNodes": {
          "description": "the minimum number of nodes that can be deployed",
          "type": "integer",
          "format": "int32"
        },
        "minNodes": {
          "description": "the minimum number of nodes that can be deployed",
          "type": "integer",
          "format": "int32",
          "minimum": 1
        },
        "name": {
          "description": "the cluster name, should be valid for use in dns names",
          "type": "string",
          "pattern": "^[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?(\\.[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?)*$",
          "readOnly": true
        },
        "noOfMasters": {
          "description": "the number of master nodes to create",
          "type": "integer",
          "format": "int32",
          "default": 1,
          "minimum": 1
        },
        "nodeResourcePools": {
          "type": "array",
          "items": {
            "type": "string",
            "minLength": 1
          }
        },
        "storageClasses": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/storageClass"
          }
        },
        "thumbprint": {
          "description": "the thumbprint of the vCenter server certificate",
          "type": "string",
          "minLength": 57,
          "pattern": "[a-fA-F0-9:]+"
        }
      }
    },
    "credentials": {
      "type": "object",
      "required": [
        "username",
        "password"
      ],
      "properties": {
        "password": {
          "type": "string",
          "format": "password",
          "minLength": 6,
          "x-nullable": false
        },
        "username": {
          "type": "string",
          "minLength": 1,
          "x-nullable": false
        }
      }
    },
    "diskFormat": {
      "type": "string",
      "default": "thin",
      "enum": [
        "thin",
        "zeroedthink",
        "eagerzeroedthick"
      ]
    },
    "error": {
      "description": "the default error model for all the error responses coming from the VCCS\n",
      "type": "object",
      "required": [
        "message",
        "code"
      ],
      "properties": {
        "cause": {
          "$ref": "#/definitions/error"
        },
        "code": {
          "description": "The error code",
          "type": "integer",
          "format": "int64"
        },
        "helpUrl": {
          "description": "link to help page explaining the error in more detail",
          "type": "string",
          "format": "uri"
        },
        "message": {
          "description": "The error message",
          "type": "string"
        }
      }
    },
    "instanceSize": {
      "description": "The sizes for an instance:\n\n* small: 1 cpu, 1Gb\n* medium: 2 cpu, 2Gb\n* large: 4 cpu, 8Gb\n* huge: 8 cpu, 32Gb\n* ginormous: 16cpu, 64Gb\n",
      "type": "string",
      "enum": [
        "small",
        "medium",
        "large",
        "huge",
        "ginormous"
      ]
    },
    "storageClass": {
      "type": "object",
      "required": [
        "name",
        "diskformat",
        "datastore"
      ],
      "properties": {
        "cacheReservation": {
          "description": "Flash read cache reservation",
          "type": "integer",
          "format": "int32"
        },
        "datastore": {
          "description": "the name of the datastore to create the volume in",
          "type": "string",
          "minLength": 1
        },
        "diskStripes": {
          "description": "Number of disk stripes per object",
          "type": "integer",
          "format": "int32"
        },
        "diskformat": {
          "allOf": [
            {
              "$ref": "#/definitions/diskFormat"
            },
            {
              "description": "the format for the disk, defaults to thin"
            }
          ]
        },
        "forceProvisioning": {
          "description": "Force provisioning",
          "type": "boolean"
        },
        "hostFailuresToTolerate": {
          "description": "Number of failures to tolerate",
          "type": "integer",
          "format": "int32"
        },
        "iopsLimit": {
          "description": "IOPS limit for object",
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "description": "the name of the storage class",
          "type": "string"
        },
        "objectSpaceReservation": {
          "description": "Object space reservation",
          "type": "integer",
          "format": "int64"
        }
      }
    },
    "task": {
      "description": "an asynchronous tasko",
      "type": "object",
      "required": [
        "id",
        "state"
      ],
      "properties": {
        "context": {
          "$ref": "#/definitions/taskContext"
        },
        "created": {
          "type": "string",
          "format": "date-time"
        },
        "id": {
          "$ref": "#/definitions/taskId"
        },
        "state": {
          "$ref": "#/definitions/taskState"
        },
        "step": {
          "$ref": "#/definitions/taskStep"
        },
        "taskType": {
          "$ref": "#/definitions/taskType"
        },
        "ttl": {
          "type": "string",
          "format": "duration"
        }
      }
    },
    "taskContext": {
      "description": "the context for a task, contains data to describe what this job pertained to.",
      "type": "object",
      "properties": {
        "cause": {
          "type": "string"
        },
        "clusterName": {
          "type": "string",
          "maxLength": 63,
          "minLength": 3,
          "pattern": "^[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?(\\.[a-zA-Z](([-0-9a-zA-Z]+)?[0-9a-zA-Z])?)*$"
        },
        "log": {
          "type": "string"
        },
        "timeout": {
          "type": "integer",
          "default": 3600
        }
      }
    },
    "taskId": {
      "type": "string",
      "minLength": 1
    },
    "taskState": {
      "type": "string",
      "enum": [
        "processing",
        "completed",
        "failed"
      ]
    },
    "taskStep": {
      "type": "string"
    },
    "taskType": {
      "type": "string",
      "enum": [
        "create",
        "delete",
        "update"
      ]
    }
  },
  "parameters": {
    "requestId": {
      "minLength": 1,
      "type": "string",
      "description": "A unique UUID for the request",
      "name": "X-Request-Id",
      "in": "header"
    }
  },
  "responses": {
    "errorClusterNameConflict": {
      "description": "The provided cluster name already exists",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorClusterNotFound": {
      "description": "The cluster was not found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorDefault": {
      "description": "Error",
      "schema": {
        "$ref": "#/definitions/error"
      }
    },
    "errorTaskNotFound": {
      "description": "The task was not found",
      "schema": {
        "$ref": "#/definitions/error"
      }
    }
  }
}`))
}
