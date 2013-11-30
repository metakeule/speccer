# speccer - a CLI tool for specs  

[HTML](http://metakeule.github.io/speccer/spec/spec.html)  
[Text (Markdown)](http://metakeule.github.io/speccer/spec/spec.md)  
[JSON](http://metakeule.github.io/speccer/spec/spec.json)  

project `speclib`  
is part of <http://metakeule.github.io/speclib/spec/spec.html>  

company `Know GmbH`  

persons  

- `Marc René Arns <linux@marcrenearns.de>` (mra)  
- `Reason why a paragraph has the current state` (reason for state)  

superseded by  

- [new](http://metakeule.github.io/speccer/spec/new)  

Translations  

- [`de_DE`](http://metakeule.github.io/speccer/spec/de_DE/spec)  

Related Specifications

- [other](http://metakeule.github.io/speccer/spec/other)  

requested by

- mra  

## Overview

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Implementing` 
| **last update**                  | `2013-11-30` 
| **deadline**                     | `soon` 
| **est. working hours**           | `2` 
| **UUID**                         | `b18c7bd2-9ffd-4150-aa67-fd1636560d98` 


[*speccer*](https://github.com/metakeule/speccer) is a CLI tool 
for dealing with specs.

It is a frontend to the [speclib](https://github.com/metakeule/speclib).



> `mra`  
>   
> just a comment

- - - -

> `reason for state`  
>   
> other comment


******
******

## Scenario


### S1 Create Specification

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-24` 
| **UUID**                         | `6d9cb892-ab9f-4236-bcb2-184aff7de019` 


Mike wants to create a specification. He gets a clean
starting point with `speccer`.

> `mra`  
>   
> comment one

- - - -

> `reason for state`  
>   
> comment2


******
******

### S2 Add comments

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `ed61b3ca-b0b3-499d-af64-308d6c06abfc` 


Mike receives comments for his specification.
He wants to add them without hassle.


******
******

### S3 Editable

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `ae2bfd27-9e91-48b2-8430-02f12741dccd` 


Mike wants to update his specification with is
editor.

******
******

### S4 Scriptability

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `f949981f-f0d0-4288-8eae-eb3fad0a9d05` 


Gill wants to change specifications with a script.

******
******

### S5 Import and export

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `3e437143-8f7f-42ea-bceb-15ff28eb40df` 


Gill wants to read import and export specifications
from different places.

******
******

### S6 View as HTML and markdown

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `d97d2941-ec37-43db-a210-29c033d52dc7` 


Sue wants to read the specifications in HTML and 
be able to mail them as markdown.


******
******

### S7 Filter for documentation

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `1df1af64-22fe-48bf-954a-6ca01aecfa84` 


Peter wants to create a documentation from the spec and
wants the unneccessary informations to be filtered.

******
******

### S8 Filter for implementation

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `743a3aa4-273b-4d89-aa3f-7ad9f466daac` 


Anne wants to implement parts of the specification.
She just wants to read the parts that are ready
to be implemented.

******
******

### S9 Templating

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `agreed` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `472220e2-848b-4629-a8cb-9bb0140cebc8` 


John wants to reuse specs. But the new spec with their parts should get 
new UUIDs to be saved in the Database.

## Non-Goal


### N1 No GUI

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `141dd94c-d5f8-4b15-aac0-c72eedcd173c` 


To have a GUI would be a project on its own.
However such a project could also be external and be built 
upon *speccer*.


******
******

### N2 No Database

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `12f5da6a-dce7-4926-ae14-1fe55d452207` 


`speccer` will operate on plain text json files only.
There will be no way to have a database and search.
A database could however exchange data with `speccer`
via json files.

Since each paragraph and therefor the spec (via `Overview` 
paragraph) has its own `UUID` it should be easy to archive
them in a database and make them searchable.

******
******

### N3 No versioning

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `f66e5709-8c6a-4605-94b0-a19561b35a0b` 


`speccer` will not do any backups or have versions.
You are advised to use a version control system to save your
specs or make backups by a backup tool.


******
******

### N4 No editing

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `e1ba4fa3-3c9d-4dc7-9c40-b72c8ea5a2b1` 


`speccer` will have any functionality of an editor.
You may however replace parts of the spec by exporting them
to your favourite editor and reimport them when you are finished.


******
******

### N5 No merging of specs

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `ec87b659-7df4-46da-a54a-853e1483a22e` 


`speccer` won't provide any mechanism to merge specs.
Since you can export and import parts of the specs in
plain text or markdown, it is not hard to merge by hand.

******
******

### N6 No rights management

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `f4523a32-adec-494c-a122-1db47f1686a7` 


`speccer` won't have any rights management baked in. 
You could manage the viewing and changing rights for
a single spec by the filesystem.

If you need finer control, you need to write a wrapper
around `speccer` or the `speclib`.

******
******

### N7 No multiple comments per author and paragraph

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `08e57c5f-aa21-4c18-936f-3fc85b8a4bdd` 


With `speclib` (and therefor `speccer` too),
a single author can only have one comment per paragraph.

We consider this a good thing, since it forces the spec
to be split up in several paragraphs if they get too many
comments. The spec will benefit from it.

Also it is possible to have comments of arbitrary length.

******
******

### N8 No 'threads'

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `51b6987d-b12a-4e94-92fd-14e8c46c9afe` 


Again this is a decision by `speclib` and considered a good 
thing. There should not be comments on comments and endless
discussions. Instead the comments should be directed to a 
paragraph and decision should be made soon.

Then the different topics should seperate into the different
sections like `Undecided`, `ContradictionS`, `DefinitionS`,
`FeatureS` etc. and there should be multiple paragraphs 
which may each have their own state of consense and comments.


******
******

### N9 No order of comments

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `34a9b1ae-bfb4-4d23-b4b1-1f07b223dc9f` 


Again this is a decision by `speclib`.
Since each author may just have one comment per paragraph, 
there is no point in have an order for comments.

However their is an order of paragraphs and that could 
be changed.


******
******

### N10 No text generation

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `b1156de4-cd51-46a0-afd5-35ada4deb2c8` 


Since the specs format is based on json, it is very easy 
and flexible to have own templates and copy them.

******
******

### N11 No PDF export

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `d39da0b6-6f52-4985-b352-003a7fcd352a` 


There are great tools to convert from HTML to PDF, like `wkhtmltopdf` and
also markdown to pdf converter like `pandoc`. Since `speccer` exports 
to HTML and markdown you should be able to use these tools to generate
PDFs.

******
******

### N12 No Translations for SPECS

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `dbe28598-cb0a-4f48-aaea-7a4caca9ce09` 


There is rudimentary support for translations of key words that
are automatically inserted via `speclib`.

If your language is not supported, it is very easy to translate 
on your own. A simple bash-script will do. You only have to
replace the keys as they are in the translations folders.

You might also contribute your translations to the project.

However for the main text of the specs they should be translated
by creating a new spec of each language and telling the other
specs of the translation in the `Translations` section of the 
specs.

******
******

### N13 No time tracking

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `b7e022fa-cadb-4458-851d-49a0cb022b3e` 


Each paragraph has a field `est. working hours` that could be used 
to do basic scheduling.

An external tool could then collect the estimated hours and 
compare them to the real time to drive conclusions about the 
schedule.

More time tracking functionality is not part of this project.


******
******

### N14 No translation of options

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `8d824161-9ca4-4a18-9c84-f716033d8737` 


In order to be able to associate the wordings of the options
to the corresponding keys in the json file and the untranslated
specs the options will not be translated.

## Contradiction


### C1 Rights Management vs responsible

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `obsolet` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `db825b93-fbbb-4442-85e3-89cd899a3699` 


**Rights Management** is a **Non Goal** but on the other hand we 
have persons responsible for paragraphs.

If anyone may edit any paragraph, then the responsibility
must constantly change (and is worthless) or it is false.


> `reason for state`  
>   
> Yes the responsability should change often in order to request feedback.  
> If multiple persons work on a spec, the spec is considered teamwork  
> and the people should work together.  
>   
> Only one person is responsible for each paragraph at a single point in  
> time. This could be enforced by external tools as well as rights management  
> if necessary. Although with version control it should not be necessary.


## Definition


### D1 inherited definitions

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `70a075a0-04ff-4692-9021-e56cb595e00e` 


All definitions of the speclib apply

******
******

### D2 current spec

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `f1930868-da24-4cb1-a26f-90c20debb975` 



The current spec is the spec that is loaded from a special file
in the current directory (`pwd`).

If the `-language` option is set, the spec is loaded from the file
`spec_[language].json`, e.g.

    spec_en_US.json

Otherwise it is looking for any files matching the regular expression

    spec*.json

If there is just one file matching, the spec is loaded from this file.
Otherwise an error is thrown mentioning that the language option should
probably be set.

******
******

### D3 name parameter

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `21431f37-09a5-42ee-a049-20d054acc849` 


The parameter passed via `-name` is the name of a spec property (like
`Translations` with the command `property`) or the name of a paragraph
property (like `state` with the command `meta`).



******
******

### D4 usage parameter

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `afbc5bfe-c0f1-48a5-8536-3f10370e4576` 


The parameter passed via `-usage` is the name of a filter provided by
`speccer` for typical usage scenarios.

There are currently:

- **documentation** filters 
  - PROPERTIES
  - META
  - CommentS
  - Non-Goal
  - Contradiction
  - Undecided
  - End of Specification
  - Files / Resources
  - planning
  - agreed
  - Implementing
  - obsolet
- **approval** filters
  - META
  - CommentS
  - Contradiction
  - Undecided
  - End of Specification
  - Files / Resources
  - planning
  - obsolet
- **discussion** filters
  - Overview
  - agreed
  - Implementing
  - obsolet
  - End of Specification
- **implementation** filters
  - Undecided
  - planning
  - Finished
  - obsolet


******
******

### D5 Linefeed export

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `cc3a6fd8-fe1d-493c-997f-5f0c0b61dcc7` 


Every import and export replaces to `\n` only. 
Editors should be able to handle it, otherwise other tools would be 
needed to convert before editing.



## Feature


### F1 create a new spec

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `7a6daf30-1bfc-4d97-8e82-dd3ae6f7380e` 


Go into a folder that should contain your spec and run

    speccer -cmd create

This creates a file `spec.json` in the current folder.

If the `-language` parameter is set (to e.g. `en_US`) it creates
the file `spec_en_US.json` instead.

Be aware that the new spec is not valid since there are missing
properties. You can validate a spec to see them:

    speccer -cmd validate





******
******

### F2 validate a spec

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `50c8c8f8-846f-4a39-81ce-23f869aa9b64` 


Go into a folder that contains your spec and run

    speccer -cmd validate

It it returns nothing your spec is valid.

******
******

### F3 export markdown to stdout

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `78af57be-97a0-4a06-8ca1-d5b2c13b25b5` 


Go into a folder that contains your spec and run

    speccer -cmd markdown

You might also pass the usage option to filter out some data, e.g.

    speccer -cmd markdown -usage "documentation"

    



******
******

### F4 export HTML to stdout

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `3f36007c-3063-4ebb-a5c3-2c9ff0bda510` 


Go into a folder that contains your spec and run

    speccer -cmd html

You might also pass the usage option to filter out some data, e.g.

    speccer -cmd html -usage "documentation"





******
******

### F5 print property of spec

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `884999ca-4e76-46c6-aeed-ef8aa89fa0e0` 


Go into a folder that contains your spec and run

    speccer -cmd property -name company

to print the `company` property. The properties are 
all written in high caps and defined in `speclib`.




******
******

### F6 set property of spec

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `e42f5105-74ee-448a-899d-68af690e6e88` 


Go into a folder that contains your spec and run

    speccer -cmd property -name company -set acme

to set the `company` property to `acme`. The properties are 
all written in high caps and defined in `speclib`.




******
******

### F7 unset property of spec

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `98f3bcf3-ea1b-4119-8219-ac33b6ad1f30` 


Go into a folder that contains your spec and run

    speccer -cmd property -name is part of -uset

to unset the `is part of` property. Only the following properties can
be unset:

- `REQUESTEDBY`
- `Related Specifications`
- `Translations`
- `SUPERSEDEDBY`
- `Files / Resources`
- `is part of`
- `persons`


******
******

### F8 add markdown text to a section

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `0e04c354-207d-4e24-a536-6d12f7db08b8` 


Go into a folder that contains your spec and write your text
into a file, let's say `new.md`.

Put the title of the paragraph in the first line 
and the body in the next lines.
Then run

    speccer -cmd add -sec Scenario -resp 'Your Name' -set new.md

to add title and body to the `Scenario` section.
The available sections are defined in the `speclib`.






******
******

### F9 show list of paragraphs

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `936f9d4a-6c84-4a3b-a496-b3022f9fc78c` 


Go into a folder that contains your spec and run

    speccer -cmd positions -sec Scenario

to see a list of all paragraphs of the `Scenario` section.

Each paragraph has the title printed preceeded by 
a number indicating the position of the paragraph inside
the section. This position can be used to print or change
the paragraph.




******
******

### F10 show content of a paragraph

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `4df2be31-ef24-4630-a8e0-587ab43b4c38` 


Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to see. Therefor you might want to first 
print all positions of the section (here `Scenario`) 
with

    speccer -cmd positions -sec Scenario

then you can run

    speccer -cmd text -sec Scenario -at 3

to see the text for the 3rd paragraph in the
`Scenario` section.

If you want to include comments and meta data, run

    speccer -cmd text -sec Scenario -at 3 -with-comments -with-meta




******
******

### F11 change content of a paragraph

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `ef854fec-9cec-4912-ae8d-6a0e2ceac098` 


Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to see. Therefor you might want to first 
print all positions of the section (here `Scenario`) 
with

    speccer -cmd positions -sec Scenario

then you can run

    speccer -cmd text -sec Scenario -at 3 > temp.md

to save the current text and comments for the 3rd paragraph in the
`Scenario` section to the file `temp.md`.

Now fire up your editor to change the content. Then run

    speccer -cmd text -sec Scenario -at 3 -set temp.md

to set the text of the 3rd paragraph in the `Scenario` section.

The comments and meta data (including the title) of the 
3rd paragraph are not modified by this operation. 
So you may want to change them too.







******
******

### F12 see, add and change comment to a paragraph

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `6daa3b1a-0cda-4861-b935-527474415efb` 


Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to see. Therefor you might want to first 
print all positions of the section (here `Scenario`) 
with

    speccer -cmd positions -sec Scenario

then you can run

    speccer -cmd comment -sec Scenario -at 3

to see the current comments for the 3rd paragraph in the
`Scenario` section.

Now if the author of your comment already made a comment to 
this paragraph and if you want to append to this old comment,
you first have to save the old one.

To do this, run 

    speccer -cmd comment -sec Scenario -at 3 -author Jim > temp.md

to save the comment of `Jim` to the file `temp.md`. Then you can edit this
file. If you just want to remove the old comment or if there was no
comment you may simply start with an empty file `temp.md`.

After you edited the file, run 

    speccer -cmd comment -sec Scenario -at 3 -author Jim -set temp.md

to set the comment of `Jim` on the 3rd paragraph of the `Scenario` section
to the content of the file `temp.md`.




******
******

### F13 move a paragraph

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `322afa61-0ddf-4743-b7b5-fbfcf953acee` 


Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to move. Therefor you might want to first 
print all positions of the section (here `Scenario`) 
with

    speccer -cmd positions -sec Scenario

then you can run

    speccer -cmd move -sec Scenario -at 3 -to 1

to move the 3rd paragraph in the `Scenario` section to be the first.




******
******

### F14 delete a paragraph

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `0c570c6c-39c5-43c2-ba60-5368260fd448` 


Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to delete. Therefor you might want to first 
print all positions of the section (here `Scenario`) 
with

    speccer -cmd positions -sec Scenario

then you can run

    speccer -cmd rm -sec Scenario -at 3

to remove the 3rd paragraph in the `Scenario` section.




******
******

### F15 see and change meta data of a paragraph

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `e9fa657c-1700-4bbe-b265-beb420eb661c` 


Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to see the meta data. Therefor you might want to first 
print all positions of the section (here `Scenario`) 
with

    speccer -cmd positions -sec Scenario

then you can run

    speccer -cmd meta -sec Scenario -at 3

to see all meta data of the 3rd paragraph in the `Scenario` section.

If you want to change a datum you will have to pass the name:

    speccer -cmd meta -sec Scenario -at 3 -name responsible person -set Jim






******
******

### F16 shortcuts

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `e3373cdc-d9c3-45ab-bc00-41d8118df8e1` 


Some commands and options are hard to remember.
But instead of bake new options into `speccer` we
deliver an alias file you could use in your bash or zsh shell.

To use the `bash` / `zsh` shortcuts, copy the `aliases.sh`
into your home directory

    cp aliases.sh ~/.speccer_aliases.sh

and write the following into your `.bashrc`

    . ~/.speccer_aliases.sh

Then you will have access to the following shortcuts
(in bold is an example input to put after the alias).

#### General
`spec_create`   
`spec_validate`  
`spec_copy`  

#### Markdown export
`spec_markdown`   
`spec_markdown-doc`   
`spec_markdown-approve`   
`spec_markdown-discuss`   
`spec_markdown-implement`   
`spec_markdown-clean`   

#### HTML export
`spec_html`   
`spec_html-doc`   
`spec_html-approve`   
`spec_html-discuss`   
`spec_html-implement`   
`spec_html-clean`   

#### Spec property getters
`spec-requestedby`   
`spec-related`   
`spec-translations`   
`spec-supersededby`   
`spec-resources`   
`spec-parent`   
`spec-company`   
`spec-project`   
`spec-persons`   
`spec-url`   
`spec-language`   
`spec-dateformat`   
`spec-approved`   

#### Spec property unsetters
`spec_uset-requestedby`   
`spec_uset-related`   
`spec_uset-translations`   
`spec_uset-supersededby`   
`spec_uset-resources`   
`spec_uset-parent`   
`spec_uset-persons`   

#### Spec property setters
`spec_set-requestedby` **Jim**  
`spec_set-related` **'{"other": "http://other.com/spec"}'**  
`spec_set-translations`  **'{"de_DE": "http://de.other.com/spec"}'**  
`spec_set-supersededby` **'{"v2": "http://other.com/v2/spec"}'**  
`spec_set-resources` **'{"ER diagramm": "http://other.com/er.png"}'**  
`spec_set-parent` **'http://other.com/parent-spec'**  
`spec_set-company` **acme**  
`spec_set-project` **'my project'**  
`spec_set-url` **'http://greatspec.com/spec'**  
`spec_set-language` **en_US**  
`spec_set-dateformat` **YYYY-MM-DD**  
`spec_set-approved`  **true**  
`spec_set-persons`  **'{"PP": "Paul Panther <paul@panther.com>"}'**  

#### Add Sections
`spec_add-Scenario`  **new.md**  
`spec_add-Contradiction`  **new.md**  
`spec_add-Definition` **new.md**  
`spec_add-Feature` **new.md**  
`spec_add-Non-Goal`  **new.md**  
`spec_add-Undecided` **new.md**  

#### Show Sections
`spec-Overview`   
`spec-Scenario-at` **2**  
`spec-Contradiction-at` **5**  
`spec-Definition-at` **20**  
`spec-Feature-at` **1**  
`spec-Non-Goal-at`  **2**  
`spec-Undecided-at`  **5**  

#### Show Sections with comments and meta
`spec-Overview-full`  
`spec-Scenario-full-at` **2**  
`spec-Contradiction-full-at` **2**  
`spec-Definition-full-at` **5**  
`spec-Feature-full-at` **2**  
`spec-Non-Goal-full-at`  **7**  
`spec-Undecided-full-at` **2**  

#### Get list of Paragraphs for all section
`spec_ls`  

#### Get list of Paragraphs for a section
`spec_ls-Scenario`   
`spec_ls-Contradiction`   
`spec_ls-Definition`   
`spec_ls-Feature`   
`spec_ls-Non-Goal`   
`spec_ls-Undecided`   

#### Get uuids of Paragraphs for a section
`spec_uuids`
`spec_uuids-Scenario`
`spec_uuids-Contradiction`
`spec_uuids-Definition`
`spec_uuids-Feature`
`spec_uuids-Non-Goal`
`spec_uuids-Undecided`

#### Get Comments of Paragraphs for a section
`spec_comment-Overview`   
`spec_comment-Scenario-at`  **2**  
`spec_comment-Contradiction-at`  **2**  
`spec_comment-Definition-at`  **2**  
`spec_comment-Feature-at`  **2**  
`spec_comment-Non-Goal-at`  **2**  
`spec_comment-Undecided-at` **2**  

#### Set Comment of Paragraph for a section
`spec_comment_replace-Overview`  **new.md**  
`spec_comment_replace-Scenario`  **new.md -at 2**  
`spec_comment_replace-Contradiction`  **new.md -at 2**  
`spec_comment_replace-Definition`  **new.md -at 2**  
`spec_comment_replace-Feature`  **new.md -at 2**  
`spec_comment_replace-Non-Goal`  **new.md -at 2**  
`spec_comment_replace-Undecided`  **new.md -at 2**  

#### Replace Sections
`spec_replace-Overview`  **new.md**  
`spec_replace-Scenario`  **new.md -at 2**  
`spec_replace-Contradiction` **new.md -at 2**  
`spec_replace-Definition` **new.md -at 2**  
`spec_replace-Feature` **new.md -at 2**  
`spec_replace-Non-Goal` **new.md -at 2**  
`spec_replace-Undecided` **new.md -at 2**  

#### Move Sections
`spec_mv-Scenario-at`  **2 -to 5**  
`spec_mv-Contradiction-at` **2 -to 5**  
`spec_mv-Definition-at` **2 -to 5**  
`spec_mv-Feature-at` **2 -to 5**  
`spec_mv-Non-Goal-at` **2 -to 5**  
`spec_mv-Undecided-at` **2 -to 5**  

#### Remove Sections
`spec_rm-Scenario-at` **2**  
`spec_rm-Contradiction-at` **2**  
`spec_rm-Definition-at` **2**  
`spec_rm-Feature-at` **2**  
`spec_rm-Non-Goal-at` **2**  
`spec_rm-Undecided-at`  **2**  

#### Get Meta Data of Overview / Spec
`spec-responsible person`  
`spec-state`  
`spec-deadline`  
`spec-last update`  
`spec-est. working hours`  
`spec-TITLE`  
`spec-UUID`

#### Get Meta Data of Section
`spec-Scenario-responsible person-at`  **2**  
`spec-Scenario-state-at`  **2**  
`spec-Scenario-deadline-at`  **2**  
`spec-Scenario-last update-at`  **2**  
`spec-Scenario-est. working hours-at`  **2**  
`spec-Scenario-TITLE-at`  **2**  
`spec-Scenario-UUID-at`  **2**  

`spec-Contradiction-responsible person-at`  **2**  
`spec-Contradiction-state-at`  **2**  
`spec-Contradiction-deadline-at`  **2**  
`spec-Contradiction-last update-at`  **2**  
`spec-Contradiction-est. working hours-at`  **2**  
`spec-Contradiction-TITLE-at`  **2**  
`spec-Contradiction-UUID-at`  **2**  

`spec-Definition-responsible person-at`  **2**  
`spec-Definition-state-at`  **2**  
`spec-Definition-deadline-at`  **2**  
`spec-Definition-last update-at`  **2**  
`spec-Definition-est. working hours-at`  **2**  
`spec-Definition-TITLE-at`  **2**  
`spec-Definition-UUID-at`  **2**  

`spec-Feature-responsible person-at`  **2**  
`spec-Feature-state-at`  **2**  
`spec-Feature-deadline-at`  **2**  
`spec-Feature-last update-at`  **2**  
`spec-Feature-est. working hours-at`  **2**  
`spec-Feature-TITLE-at`  **2**  
`spec-Feature-UUID-at`  **2**  

`spec-Non-Goal-responsible person-at`  **2**  
`spec-Non-Goal-state-at`  **2**  
`spec-Non-Goal-deadline-at`   **2**  
`spec-Non-Goal-last update-at`  **2**  
`spec-Non-Goal-est. working hours-at`  **2**  
`spec-Non-Goal-TITLE-at`  **2**  
`spec-Non-Goal-UUID-at`  **2**  

`spec-Undecided-responsible person-at`  **2**  
`spec-Undecided-state-at`  **2**  
`spec-Undecided-deadline-at`  **2**  
`spec-Undecided-last update-at`  **2**  
`spec-Undecided-est. working hours-at` **2**  
`spec-Undecided-TITLE-at`  **2**  
`spec-Non-Goal-UUID-at`  **2**  

#### Set Meta Data of Overview / Spec
`spec_set-responsible person` **'Peter Pan'**  
`spec_set-planning`  
`spec_set-agreed`  
`spec_set-Implementing`  
`spec_set-Finished`   
`spec_set-obsolet`  
`spec_set-deadline`  **'2014-12-24'**  
`spec_set-est. working hours`  **12**  
`spec_set-TITLE`  **'my title'**  

#### Set Meta Data of Section

`spec_set-Scenario-responsible person`   **'Peter Pan' -at 2**    
`spec_set-Scenario-planning`  **-at 2**  
`spec_set-Scenario-agreed`   **-at 2**  
`spec_set-Scenario-Implementing`   **-at 2**  
`spec_set-Scenario-Finished`   **-at 2**  
`spec_set-Scenario-obsolet`   **-at 2**  
`spec_set-Scenario-deadline`   **'2014-12-24' -at 2**  
`spec_set-Scenario-est. working hours`  **12 -at 2**  
`spec_set-Scenario-TITLE`  **'my title' -at 3**  

`spec_set-Contradiction-responsible person`  **'Peter Pan' -at 2**    
`spec_set-Contradiction-planning`   **-at 2**  
`spec_set-Contradiction-agreed`   **-at 2**  
`spec_set-Contradiction-Implementing`   **-at 2**  
`spec_set-Contradiction-Finished`   **-at 2**  
`spec_set-Contradiction-obsolet`   **-at 2**  
`spec_set-Contradiction-deadline`  **'2014-12-24' -at 2**  
`spec_set-Contradiction-est. working hours`  **3 -at 2**  
`spec_set-Contradiction-TITLE`  **'my title' -at 3**  

`spec_set-Definition-responsible person` **'Peter Pan' -at 2**    
`spec_set-Definition-planning`   **-at 2**  
`spec_set-Definition-agreed`   **-at 2**  
`spec_set-Definition-Implementing`   **-at 2**  
`spec_set-Definition-Finished`   **-at 2**  
`spec_set-Definition-obsolet`   **-at 2**  
`spec_set-Definition-deadline`  **'2014-12-24' -at 2**  
`spec_set-Definition-est. working hours`  **3 -at 2**  
`spec_set-Definition-TITLE`  **'my title' -at 3**  

`spec_set-Feature-responsible person`  **'Peter Pan' -at 2**    
`spec_set-Feature-planning`  **-at 2**  
`spec_set-Feature-agreed`  **-at 2**  
`spec_set-Feature-Implementing`  **-at 2**  
`spec_set-Feature-Finished`  **-at 2**  
`spec_set-Feature-obsolet`   **-at 2**  
`spec_set-Feature-deadline`  **'2014-12-24' -at 2**  
`spec_set-Feature-est. working hours`  **3 -at 2**  
`spec_set-Feature-TITLE`  **'my title' -at 3**  

`spec_set-Non-Goal-responsible person` **'Peter Pan' -at 2**    
`spec_set-Non-Goal-planning`   **-at 2**  
`spec_set-Non-Goal-agreed`   **-at 2**  
`spec_set-Non-Goal-Implementing`  **-at 2**  
`spec_set-Non-Goal-Finished`   **-at 2**  
`spec_set-Non-Goal-obsolet`   **-at 2**  
`spec_set-Non-Goal-deadline`  **'2014-12-24' -at 2**  
`spec_set-Non-Goal-est. working hours`  **3 -at 2**  
`spec_set-Non-Goal-TITLE`  **'my title' -at 3**  

`spec_set-Undecided-responsible person`  **'Peter Pan' -at 2**    
`spec_set-Undecided-planning`   **-at 2**  
`spec_set-Undecided-agreed`   **-at 2**  
`spec_set-Undecided-Implementing`   **-at 2**  
`spec_set-Undecided-Finished`   **-at 2**  
`spec_set-Undecided-obsolet`   **-at 2**  
`spec_set-Undecided-deadline`  **'2014-12-24' -at 2**  
`spec_set-Undecided-est. working hours`  **3 -at 2**  
`spec_set-Undecided-TITLE`  **'my title' -at 3**  



******
******

### F17 adding multiple paragraphs to a section in one go

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `planning` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `98aa38cd-1fd4-49ff-88ae-6c5f42e0ec89` 


Go into a folder that contains your spec.

Create a new file (e.g. `new.md`) and write your paragraphs
into this file.

Each paragraph should hold its title inside its first line
and should end with a line containing `***`.

Then run

    speccer -cmd add-multi -sec Scenario -resp Jim -set new.md

to add multiple paragraphs to the Scenario section as defined in
`new.md` with all having the responsible set to Jim.



******
******

### F18 fix breakage caused by changes in the speclib

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-24` 
| **est. working hours**           | `2` 
| **UUID**                         | `ad6af27a-b991-488e-9d56-09bf33f181e3` 


The `speclib` changed in incompatible ways. This affects also 
the format of the **json** file.

### TODO

- fix speccer to conform to the new API
- manually convert the existing `spec.json` of speccer 
  to the new format

******
******

### F19 setting multiple comments to a paragraph in one go

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `planning` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `ccc8d79d-6618-44c7-878e-72e295301e8b` 


Go into a folder that contains your spec.

Create a new file (e.g. `new.md`) and write the comments
into this file.

Each comment should hold its author inside the first line
and should end with a line containing `***`

You will need to know the position of the paragraph
you want to set the comments to. Therefor you might want to first 
print all positions of the section (here `Scenario`) 
with

    speccer -cmd positions -sec Scenario

then you can run

    speccer -cmd comment-multi -sec Scenario -at 3 -set new.md

to set multiple comments to the 3rd paragraph of the `Scenario` 
section as defined in `new.md`.

    

******
******

### F20 change states

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-24` 
| **est. working hours**           | `3` 
| **UUID**                         | `bdb1bb6e-6c96-4138-b896-4c678466c767` 


some states make no sense for some sections, e.g.

- `Non-Goal` section can't be "implemented" in any 
   meaningful way

In think the following states should change the name:

- `approved`           => `agreed`
- `PARTLY_IMPLEMENTED` => `Implementing`
- `FULLY_IMPLEMENTED`  => `Finished`



******
******

### F21 packing multiple paragraphs to one single paragraph

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `planning` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `5f326a27-19b8-42c1-b98a-7589e309c0b7` 


Go into a folder that contains your spec.

You will need to know the positions of the paragraph you want
to pack. Therefor you might want to first 
print all positions of the section (here `Scenario`) 
with

    speccer -cmd positions -sec Scenario

then you can run

    speccer -cmd pack -sec Scenario -packing '[3,4,5]' -resp Jim -title 'all together'

to pack the paragraphs no 3,4 and 5 of the `Scenario` section into a
new paragraph with the title `all together` and the responsible `Jim`.
The previous titles are inserted as subheadings and the previous paragraphs
are deleted.


******
******

### F22 move paragraph from one section to another

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `agreed` 
| **last update**                  | `2013-11-23` 
| **UUID**                         | `bb295d1a-4609-44e8-84fb-d42486e2f9fa` 


Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to move. Therefor you might want to first 
print all positions of the section (here `Scenario`) 
with

    speccer -cmd positions -sec Scenario

then you can run

    speccer -cmd setsec -sec Scenario -at 3 -tosec Non-Goal

to move the 3rd paragraph in the `Scenario` section 
to the `Non-Goal` section.




******
******

### F23 get the uuids of a section

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `Finished` 
| **last update**                  | `2013-11-24` 
| **UUID**                         | `a1af5f0a-b5a3-4fef-a6c0-10a91bea14ec` 


Go into a folder that contains your spec and run

    speccer -cmd uuids -sec Scenario

to see a list of all paragraphs of the `Scenario` section.

Each paragraph has the position and the uuid followed by 
the title of the paragraph.




******
******

### F24 copy spec and create new uuids for every paragraph

| **property**                     | **value**            | 
| :------------------------------- | :------------------- | 
| **responsible person**           | `mra` 
| **state**                        | `agreed` 
| **last update**                  | `2013-11-29` 
| **UUID**                         | `7e686080-17ba-4885-bcb9-f5de5da2c227` 


Go into a folder that contains the spec you want to copy and run

    speccer -cmd copy > target/spec.json

This creates a file `spec.json` in the folder `target` as copy
of the current spec. Inside the copy every `UUID` is replaced 
by a new one and all comments are removed.

Be aware that the copy may not be valid since there may be missing
properties. You can validate the new copy like this:

    cd target && speccer -cmd validate



## Files / Resources

- [img](http://metakeule.github.io/speccer/spec/spec.png)  

******
******
******
******


