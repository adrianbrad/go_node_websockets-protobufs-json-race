Current outcome from the ./results/results.txt. Also check the pdfs in the results package for the cpu profiling.

####Messages used:
#####Small:
```json
{
    "integer":9,
    "float":1.1
}
```

#####Medium:
```json
{
    "integer":2147483647,
    "float":3.4028234663852886e+38,
    "pairs": [
        {
            "text":"Lorem ipsum",
            "number":2147483647
        }
    ]
}
```

#####Big:
```json
{  
   "integer":9223372036854775807,
   "float":1.7976931348623157e+308,
   "pairs":[  
      {  
         "text":"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
         "number":9223372036854775807
      }
   ]
}
```


```
Small messages
100000 times
Protobufs elapsed time: 4.820889089s
JSON elapsed time: 5.254548231s

Medium messages
100000 times
Protobufs elapsed time: 5.34853504s
JSON elapsed time: 5.747886991s

Big messages
100000 times
Protobufs elapsed time: 5.335504258s
JSON elapsed time: 6.457228259s
```