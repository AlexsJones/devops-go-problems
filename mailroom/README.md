# Mailroom

## Problem Definition
The international mail organisation is looking at improving their mail delivery to Australia. Each letter needs to have a postcode, street address, and state to ensure that it can be delivered.  
Any post that is missing any of these will be rejected and returned back to sender. In addition, they want to gather statistics as to how many letters are sent to each state at the end of the day.

### stage 1
- [ ] Given a json file defined of `[{state, postcode, address}, ...]`
  - With each value in the json file stored as strings
  - When you go to Marshal the json object to the [post struct](solution/post.go)
  - Then the data should be correctly stored inside the struct
  - And be obtainable via the methods applied.

- [ ] Given that you have a [post](solution/post.go)
  - That implements the interface [letter](types/letter.go)
  - Then the validate method should return true
  - When postcode is valid for the state

### stage 2
 - [ ] Given that you have loaded all letters from a file
   - When processing each letter
   - Then each letter that is valid should update the correct count for the state
   - And any invalid letters should be set aside for later

 - [ ] Given that a Mailroom has processed letters
   - When I ask for the delivery count for a State
   - Then it should return the count of letters that were "sent" to that state.

### Postcode ranges
| State/Territory              | Abbreviation | Postcode range      |
| ---------------------------- | ------------ | ------------------- |
| New South Wales              | NSW          | 1000—2999           |
| Victoria                     | VIC          | 3000—3999 8000—8999 |
| Queensland                   | QLD          | 4000—4999 9000—9999 |
| South Australia              | SA           | 5000—5999           |
| Western Australia            | WA           | 6000—6999           |
| Tasmania                     | TAS          | 7000—7999           |


## Where to get started
In this problem, we have created some boiler plate code to help assist with problem
however some methods have been left blank.
Within the [types folder](types/) is all the interfaces that you are required to implement within the [solution folder](solution).

## Testing your code
We have written the solution folder to ensure that you have implemented a solution to our problem correctly.
```sh
# To run all of our test cases
go test -v ./...
```
