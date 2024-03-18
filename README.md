# spg (Static Page Generator)
Hugo alternative but much more minimal and simpler to understand. The goal is to create a command line utility that amplifies the html language just a little bit more with `html/template.go` and besides gives you the ability to get content from `markdown` files.

No:
- Templates
- Org-mode support
- Config files
- ...

## Directory structure
- content/
- layout/
- public/
- static/

### public/
Generated static pages are put here.

### static/
Directory for files that always get copied into public directory. For html files some `Layout functions` can be used.

### content/
Directory for data in transform step `data+layout=pages`. Here go directories and `.md` files.
#### .md file
Md file example:
```md
---
title: Cordon bleu
time: 60min
---

## Step 1
You need to buy required ingredients first: ..
```

Md file consists of head in yaml format surrounded with `---` and Markdown content for a page.

### layout/
Directory for layout in transform step `data+layout=pages`. Here go directories and `.html` files.
#### Intro
In the transformation step for each content file (and dir) **one content file** is matched with **one layout file** (in the dir case **nothing** is matched with **one layout file**). If the layout match is not found a page for content is not generated. The layout matching algorithm follows the procedure:
```
+--------------+                                                               
|Transform step|                                                               
++-------------+                                                               
 |                                                                             
 |   +--------------------------------+                                        
 +--->for each ELEMENT inside content/|                                        
     ++-------------------------------+                                        
      |                                                                        
     +v----------------------------------------------------------------------+ 
     |                     Search for perfect path match                     | 
     |    ex. for files (content/recepie/cb.md and layout/recepie/cb.html)   | 
     |     ex. for directories (content/recepie/ and layout/recepie.html)    | 
     |(file extension does not count since .md and .html are expected anyway)| 
     ++--------------------------+-------------------------------------------+ 
      |                          |                                             
      v                          v                                             
      If not found               If found                                      
      |                          |                                             
      |                          |                  +-------------------+      
      |                          +------------>MATCH|Identical path name|      
      |                                             |wo. file extension |      
      |                                             +-------------------+      
      |                                                                        
     +v-----------------+                           +----------------------+   
     |if ELEMENT is FILE+--------------------->MATCH|layout/magic/file.html|   
     ++-----------------+                           +----------------------+   
      |                                                                        
      v                                                                        
      else                                                                     
      |                                                                        
     +v----------------+                            +---------------------+    
     |if ELEMENT is DIR+---------------------->MATCH|layout/magic/dir.html|    
     +-----------------+                            +---------------------+    
                                                                               
+----------------------------------------+                                     
|After Transform step for yaml lists:    |                                     
|---                                     |                                     
|listname: [element, element2, element3] |                                     
|---                                     |                                     
++---------------------------------------+                                     
 |                                                                             
 |   +------------------+                 +---------------------------+        
 +--->for each yaml LIST+----------->MATCH|layout/magic/list/LIST.html|        
 |   +------------------+                 +---------------------------+        
 |                                                                             
 |   +--------------------------+         +-----------------------------------+
 +--->for each yaml LIST ELEMENT+--->MATCH|layout/magic/list/LIST/ELEMENT.html|
     +--------------------------+         +-----------------------------------+
```
#### layout/magic/
here are files with magic functions that are used in transformation step if perfect match is not found.
- `layout/magic/file.html`: matches any file
- `layout/magic/DIR/file.html`: matches any file in `DIR` directory. This has higher precedence than the parent `file.html` matchers.
- `layout/magic/dir.html`: matches any dir
- `layout/magic/DIR/dir.html`: matches any dir in `DIR` directory. This has higher precedence than the parent `dir.html` matchers.
- `layout/magic/list/*`: list matching described in above diagram
#### Layout Language
#### Layout Functions
