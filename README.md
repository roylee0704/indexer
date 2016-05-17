# indexer
Indexer, an experimental project in Golang.


## Functional Requirements
1. **Parse** a file!
  - text file at the moment.

2. **Split** strings of words!
  - splitFunc(buf []byte, atEOF bool) (advance int, token []byte, err)
  - build hit-item for each word.
  - build word-frequency.
  - goal: fast building construction.

3. **Index** it!
  - persistent it.
  - goal: fast insertion.


---

## Specification:

### SplitFunc
This is under an assumption that a token is surrounded by control-breaks. To obtain the token, ignore first-half of control-breaks, and a token is found when last-half of control-break found. i.e: "cbcbcbcb**TokenFound**cbcb"

In general, there are 2 cases in the function:

1. control-break found: return index(end-of-term), token, nil.
2. control-break !found:
  - !atEOF:  request for more data.
  - atEOF: return len(token), token, err = finalToken

####Example: ScanWords(control-break = '/space')
**Case#1**: "ABC ".
**Case#2**: "ABC".
