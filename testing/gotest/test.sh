#!/bin/bash
# Dummy TEST file
# Demostrates the use of various functions
# CANNOT use the install | instantiate | invoke | query

# 1. Include the utility script
source  utest.sh

# 2. Test case
set_test_case   'Dummy Test Case #1'
print_info      'This is an information message'
print_info      'This is another test message'
assert_equal    "10"  "10"

# 3. Test case 
set_test_case   'Dummy Test Case #2'
print_failure   'This is just a failure message'
assert_number   "300" "-le" "200"   

# 3. Test case 
set_test_case   'Dummy Test Case #3 - shows use of assert_boolean'
print_info   'This is showing the use of assert_boolean '
assert_boolean   true 