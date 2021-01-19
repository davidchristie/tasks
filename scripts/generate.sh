#!/bin/bash

# Backend code generation
(go generate ./...)

# Frontend code generation
(cd web && npm run generate)
