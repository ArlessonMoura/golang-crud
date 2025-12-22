#!/bin/bash

# Colors for output
GREEN='\033[0;32m'
RED='\033[0;31m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

BASE_URL="http://localhost:8080/api"

echo -e "${YELLOW}========== ORGANIZATION API TESTS ==========${NC}\n"

# Test 1: Create Organization
echo -e "${YELLOW}[TEST 1] Creating organization...${NC}"
ORG_RESPONSE=$(curl -s -X POST "$BASE_URL/org" \
  -H "Content-Type: application/json" \
  -d '{"name": "Tech Company"}')
echo "Response: $ORG_RESPONSE"
ORG_ID=$(echo $ORG_RESPONSE | grep -o '"id":[0-9]*' | grep -o '[0-9]*')
echo -e "${GREEN}Organization ID: $ORG_ID${NC}\n"

# Test 2: Create another Organization
echo -e "${YELLOW}[TEST 2] Creating another organization...${NC}"
ORG_RESPONSE2=$(curl -s -X POST "$BASE_URL/org" \
  -H "Content-Type: application/json" \
  -d '{"name": "StartUp Inc"}')
echo "Response: $ORG_RESPONSE2"
ORG_ID2=$(echo $ORG_RESPONSE2 | grep -o '"id":[0-9]*' | grep -o '[0-9]*')
echo -e "${GREEN}Organization ID: $ORG_ID2${NC}\n"

# Test 3: List Organizations
echo -e "${YELLOW}[TEST 3] Listing all organizations...${NC}"
curl -s -X GET "$BASE_URL/org" \
  -H "Content-Type: application/json" | jq '.'
echo ""

# Test 4: Get Organization by ID
echo -e "${YELLOW}[TEST 4] Getting organization by ID...${NC}"
curl -s -X GET "$BASE_URL/org/$ORG_ID" \
  -H "Content-Type: application/json" | jq '.'
echo ""

# Test 5: Update Organization
echo -e "${YELLOW}[TEST 5] Updating organization...${NC}"
curl -s -X PUT "$BASE_URL/org/$ORG_ID" \
  -H "Content-Type: application/json" \
  -d '{"name": "Tech Company Updated"}' | jq '.'
echo ""

# Test 6: Create Users for testing
echo -e "${YELLOW}[TEST 6] Creating test users...${NC}"
USER1=$(curl -s -X POST "$BASE_URL/users" \
  -H "Content-Type: application/json" \
  -d '{"name": "John Doe", "email": "john@example.com"}')
USER_ID1=$(echo $USER1 | grep -o '"id":[0-9]*' | grep -o '[0-9]*')
echo -e "${GREEN}User 1 ID: $USER_ID1${NC}"

USER2=$(curl -s -X POST "$BASE_URL/users" \
  -H "Content-Type: application/json" \
  -d '{"name": "Jane Smith", "email": "jane@example.com"}')
USER_ID2=$(echo $USER2 | grep -o '"id":[0-9]*' | grep -o '[0-9]*')
echo -e "${GREEN}User 2 ID: $USER_ID2${NC}\n"

# Test 6a: Get User by ID (User 1)
echo -e "${YELLOW}[TEST 6a] Getting user 1 by ID...${NC}"
curl -s -X GET "$BASE_URL/users/$USER_ID1" \
  -H "Content-Type: application/json" | jq '.'
echo ""

# Test 6b: Get User by ID (User 2)
echo -e "${YELLOW}[TEST 6b] Getting user 2 by ID...${NC}"
curl -s -X GET "$BASE_URL/users/$USER_ID2" \
  -H "Content-Type: application/json" | jq '.'
echo ""

# Test 7: Add User to Organization
echo -e "${YELLOW}[TEST 7] Adding users to organization...${NC}"
curl -s -X POST "$BASE_URL/org/$ORG_ID/users" \
  -H "Content-Type: application/json" \
  -d "{\"user_id\": $USER_ID1, \"permission\": \"ROOT\"}" | jq '.'
echo ""

curl -s -X POST "$BASE_URL/org/$ORG_ID/users" \
  -H "Content-Type: application/json" \
  -d "{\"user_id\": $USER_ID2, \"permission\": \"WRITE\"}" | jq '.'
echo ""

# Test 8: List Organization Users
echo -e "${YELLOW}[TEST 8] Listing organization users...${NC}"
curl -s -X GET "$BASE_URL/org/$ORG_ID/users" \
  -H "Content-Type: application/json" | jq '.'
echo ""

# Test 9: Update User Permission
echo -e "${YELLOW}[TEST 9] Updating user permission in organization...${NC}"
curl -s -X PUT "$BASE_URL/org/$ORG_ID/users/$USER_ID2" \
  -H "Content-Type: application/json" \
  -d '{"permission": "READ"}' | jq '.'
echo ""

# Test 10: Remove User from Organization
echo -e "${YELLOW}[TEST 10] Removing user from organization...${NC}"
curl -s -X DELETE "$BASE_URL/org/$ORG_ID/users/$USER_ID2" \
  -H "Content-Type: application/json" | jq '.'
echo ""

# Test 11: List Organization Users again
echo -e "${YELLOW}[TEST 11] Listing organization users after removal...${NC}"
curl -s -X GET "$BASE_URL/org/$ORG_ID/users" \
  -H "Content-Type: application/json" | jq '.'
echo ""

# Test 12: Delete Organization
echo -e "${YELLOW}[TEST 12] Deleting organization...${NC}"
curl -s -X DELETE "$BASE_URL/org/$ORG_ID" \
  -H "Content-Type: application/json" | jq '.'
echo ""

# Test 13: List Organizations after deletion
echo -e "${YELLOW}[TEST 13] Listing organizations after deletion...${NC}"
curl -s -X GET "$BASE_URL/org" \
  -H "Content-Type: application/json" | jq '.'
echo ""

echo -e "${GREEN}========== TESTS COMPLETED ==========${NC}"
