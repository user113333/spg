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
- static/

### static/
Directory for files that always get copied into output directory. Transformation is applied to them (it matches **nothing** with **this file**), so you get access to Go Template functions.

### content/
Directory consisting of content files for layout. A content file must end in `.md` and be written in **MarkDown**. The top of a MarkDown file may be used for content data section written in **YAML**. Section starts and ends with three dashes: `---`.

Md file example:
```md
---
title: Carbonara
time: 60min
tags: ["breakfast", "launch", "dinner"]
---

## Step 1
You need to buy required ingredients first: ..
```

### layout/
Directory consisting of layout files for describing a content transformation.
#### Intro
In the transformation step for each content file it tries to match **one content file** with **one layout file** and for each directory name it tries to match **one directory** with **one layout file**. If the match happens content gets transformed and you get some output file but if no match is found it tries to match with **magic** files inside `layout/magic/` directory and if it is unable to find a match is simply doesn't output anything - It just skips the file.

The layout matching algorithm for content files follows the following procedure:
1. First search for perfect path match from content, layout forward.
    - Ex: `content/recipes/carbonara.md` matches `layout/recipes/carbonara.html`
    - Ex: `content/recipes` matches layout/recipes.html (recipes is directory)
2. If not found, try to match with `layout/magic/file.html` (if file) or `layout/magic/dir.html` (if is directory)
3. If not found, skip
This is for matching files (and directories) from `content/` folder. In each content file a YAML section can have a sequence of items. Similarly for each sequence it tries to match **one sequence** with **one layout file** and **one sequence item** with **one layout file**. Of course there are many content files and each content file can have a same sequence name - The program just joins the "sequence names" together and it tries to match for each unique sequence name. The items for each "sequence name" are presented together for layout file. And vice versa.

The layout matching algorithm for sequences follows the procedure:
1. First search for perfect path match.
    - Ex: in sequence `tags: ["i1", "i2"]` it matches with `layout/sequence/tags.html`
2. If not found, try to match with `layout/magic/sequence.html`
3. If not found, skip

The layout matching algorithm for sequence items follows the procedure:
1. First search for perfect path match.
    - Ex: in sequence `tags: ["i1", "i2"]` it matches with `layout/sequence/tags/i1.html`
2. If not found, try to match with `layout/sequence/SEQUENCE/item.html`
2. If not found, try to match with `layout/magic/item.html`
3. If not found, skip

Note that `content/` can't have a directory with a name `magic` and `sequence`[^1].
[^1]: If this is a huge deal-breaker it can be changed so for perfect content file path match it looks inside of `layout/content/*` instead of `layout/*`.
#### layout/magic/
List of magic files:
- `layout/magic/file.html`: matches any file
- `layout/magic/DIR/file.html`: matches any file in `DIR` directory. This has higher precedence than the parent `file.html` matchers.
- `layout/magic/dir.html`: matches any directory
- `layout/magic/DIR/dir.html`: matches any directory in `DIR` directory. This has higher precedence than the parent `dir.html` matchers.
- `layout/magic/sequence.html` matches any sequence
- `layout/magic/item.html` matches any sequence item
#### Layout Language
Go template: data-driven templates for generating textual output. **TODO**
#### Go template Functions
List of Go template functions: **TODO**
