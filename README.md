# shed

```
                hhhhhhh                                             d::::::d
                h:::::h                                             d::::::d
                h:::::h                                             d::::::d
                h:::::h                                             d:::::d
    ssssssssss   h::::h hhhhh           eeeeeeeeeeee        ddddddddd:::::d
  ss::::::::::s  h::::hh:::::hhh      ee::::::::::::ee    dd::::::::::::::d
ss:::::::::::::s h::::::::::::::hh   e::::::eeeee:::::ee d::::::::::::::::d
s::::::ssss:::::sh:::::::hhh::::::h e::::::e     e:::::ed:::::::ddddd:::::d
 s:::::s  ssssss h::::::h   h::::::he:::::::eeeee::::::ed::::::d    d:::::d
   s::::::s      h:::::h     h:::::he:::::::::::::::::e d:::::d     d:::::d
      s::::::s   h:::::h     h:::::he::::::eeeeeeeeeee  d:::::d     d:::::d
ssssss   s:::::s h:::::h     h:::::he:::::::e           d:::::d     d:::::d
s:::::ssss::::::sh:::::h     h:::::he::::::::e          d::::::ddddd::::::dd
s::::::::::::::s h:::::h     h:::::h e::::::::eeeeeeee   d:::::::::::::::::d
 s:::::::::::ss  h:::::h     h:::::h  ee:::::::::::::e    d:::::::::ddd::::d
  sssssssssss    hhhhhhh     hhhhhhh    eeeeeeeeeeeeee     ddddddddd   ddddd
  ```
  
  Shell script scheduling.
  
  Automate repetitive tasks.
  
  Hooya.
  
  ```
  >>> push cd ~/Work
Pushing -> cd
>>> push ls -la
Pushing -> ls
>>> list
+------+------------------+-----------+----------+------------+
| STEP | RESOURCE LOCATOR | VALIDATED | EXECUTED | SUCCESSFUL |
+------+------------------+-----------+----------+------------+
|    1 | cd               | ✓         | ✗        | ?          |
|    2 | ls               | ✓         | ✗        | ?          |
+------+------------------+-----------+----------+------------+
>>> run
+------+------------------+-----------+----------+------------+
| STEP | RESOURCE LOCATOR | VALIDATED | EXECUTED | SUCCESSFUL |
+------+------------------+-----------+----------+------------+
|    1 | cd               | ✓         | ✓        | ✓          |
+------+------------------+-----------+----------+------------+
```
