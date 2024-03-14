# spg (Static Page Generator)

## Directory structure

- content/
- layout/
- public/
- static/

### public/
Generated static pages are put here.

### static/
Directory for files that always get copied into public directory. A file in layout is copied `n` times while static file is copied `1` time.

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
┌────────────────────────────────┐                                       
│for each ELEMENT inside content/│                                       
└┬───────────────────────────────┘                                       
 │                                                                       
┌▼──────────────────────────────────────────────────────────────────────┐
│                     Search for perfect path match                     │
│    ex. for files (content/recepie/cb.md and layout/recepie/cb.html)   │
│     ex. for directories (content/recepie/ and layout/recepie.html)    │
│(file extension does not count since .md and .html are expected anyway)│
└┬──────────────────────────┬───────────────────────────────────────────┘
 │                          │                                            
 ▼                          ▼                                            
 If not found               If found                                     
 │                          │                                            
 │                          │                  ┌───────────────────┐     
 │                          └────────────►MATCH│Identical path name│     
 │                                             │wo. file extension │     
 │                                             └───────────────────┘     
 │                                                                       
┌▼─────────────────┐                           ┌──────────────────────┐  
│if ELEMENT is FILE├─────────────────────►MATCH│layout/magic/file.html│  
└┬─────────────────┘                           └──────────────────────┘  
 │                                                                       
 ▼                                                                       
 else                                                                    
 │                                                                       
┌▼────────────────┐                            ┌─────────────────────┐   
│if ELEMENT is DIR├──────────────────────►MATCH│layout/magic/dir.html│   
└─────────────────┘                            └─────────────────────┘   
                                                                         
┌────────────────────────────────────────┐                               
│After transformation step for yaml list:│                               
│---                                     │                               
│listname: [element, element2, element3] │                               
│---                                     │                               
└────────────────────────────────────────┘                               
                                                                         
┌──────────────────┐                ┌───────────────────────────┐        
│for each yaml LIST├──────────►MATCH│layout/magic/list/LIST.html│        
└──────────────────┘                └───────────────────────────┘        
                                                                         
┌──────────────────────────┐        ┌───────────────────────────────────┐
│for each yaml LIST ELEMENT├──►MATCH│layout/magic/list/LIST/ELEMENT.html│
└──────────────────────────┘        └───────────────────────────────────┘
```
#### layout/special/
here are files with special functions that are used in transformation step if perfect match is not found.
- page.
#### Layout Language
