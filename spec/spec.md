# speccer - a CLI tool for specs  
[HTML](http://rawgithub.com/metakeule/speccer/master/spec/spec.html)  
[Text (Markdown)](http://rawgithub.com/metakeule/speccer/master/spec/spec.md)  
[JSON](http://rawgithub.com/metakeule/speccer/master/spec/spec.json)  

PARENT <http://rawgithub.com/metakeule/speclib/master/spec/spec.html>  

PROJECT `speclib`  
COMPANY `Know GmbH`  

PERSONS  

- `Marc René Arns <linux@marcrenearns.de>` (mra)  

REQUESTED_BY

- mra  

## OVERVIEW

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-11` ESTIMATEDHOURS `0`  

[*speccer*](https://github.com/metakeule/speccer) is a CLI tool for dealing with specs.

It is a frontend to the [speclib](https://github.com/metakeule/speclib).




## SCENARIOS

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### Create Specification

Mike wants to create a specification. He gets a clean
starting point with `speccer`.

******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### Add comments

Mike receives comments for his specification.
He wants to add them without hassle.


******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### See and update specification

Mike wants to update his specification with is
editor.


******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### Script

Gill wants to change specifications with a script.


******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### Import and export

Gill wants to read import and export specifications
from different places.


******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### View as HTML and markdown

Sue wants to read the specifications in HTML and 
be able to mail them as markdown.


******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### Filter for documentation

Peter wants to create a documentation from the spec and
wants the unneccessary informations to be filtered.

******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### Filter for implementation

Anne wants to implement parts of the specification.
She just wants to read the parts that are ready
to be implemented.

## NONGOALS

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No GUI

To have a GUI would be a project on its own.
However such a project could also be external and be built upon speccer.

******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No Database

`speccer` will operate on plain text json files only.
There will be no way to have a database and search.
A database could however exchange data with `speccer`
via json files.

******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No versioning

`speccer` will not do any backups or have versions.
You are advised to use a version control system to save your
specs or make backups by a backup tool.

******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No editing

`speccer` will have any functionality of an editor.
You may however replace parts of the spec by exporting them
to your favourite editor and reimport them when you are finished.

******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No merging of specs

`speccer` won't provide any mechanism to merge specs.
Since you can export and import parts of the specs in
plain text or markdown, it is not hard to merge by hand.

******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No copying of specs

`speccer` won't help you copying specs, since you can simply
copy the json files yourself and then edit the copy.

******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No rights management

`speccer` won't have any rights management baked in. 
You could manage the viewing and changing rights for
a single spec by the filesystem.

If you need finer control, you need to write a wrapper
around `speccer` or the `speclib`.

******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No multiple comments per author and paragraph

With `speclib` (and therefor `speccer` too),
a single author can only have one comment per paragraph.

We consider this a good thing, since it forces the spec
to be split up in several paragraphs if they get too many
comments. The spec will benefit from it.

Also it is possible to have comments of arbitrary length.


******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No 'threads'

Again this is a decision by `speclib` and considered a good 
thing. There should not be comments on comments and endless
discussions. Instead the comments should be directed to a 
paragraph and decision should be made soon.

Then the different topics should seperate into the different
sections like 'UNDECIDED', 'CONTRADICTIONS', 'DEFINITIONS',
'FEATURES' etc. and there should be multiple paragraphs 
which may each have their own state of consense and comments.


******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No order of comments

Again this is a decision by `speclib`.
Since each author may just have one comment per paragraph, 
there is no point in have an order for comments.

However their is an order of paragraphs and that could 
be changed.


******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No text generation

Since the specs format is based on json, it is very easy 
and flexible to have own templates and copy them.

******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No PDF export

There are great tools to convert from HTML to PDF, like `wkhtmltopdf` and
also markdown to pdf converter like `pandoc`. Since `speccer` exports 
to HTML and markdown you should be able to use these tools to generate
PDFs.

******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No Translations for SPECS

There is rudimentary support for translations of key words that
are automatically inserted via `speclib`.

If your language is not supported, it is very easy to translate 
on your own. A simple bash-script will do. You only have to
replace the keys as they are in the translations folders.

You might also contribute your translations to the project.

However for the main text of the specs they should be translated
by creating a new spec of each language and telling the other
specs of the translation in the TRANSLATIONS section of the specs.

******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No time tracking

Each paragraph has a field ESTIMATEDHOURS that could be used to
do basic scheduling.

An external tool could then collect the estimated hours and compare
them to the real time to drive conclusions about the schedule.

More time tracking functionality is not part of this project.


******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No shortcuts for commands or presets

Some commands and options are hard to remember.
But instead of bake new options into `speccer` we
deliver an alias file you could use in your bash or zsh shell.

Have a look at it to see how easy you could do one on your own.
In the alias you could also set defaults for responsibles etc.


******

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### No translation of options

In order to be able to associate the wordings of the options
to the corresponding keys in the json file and the untranslated
specs the options will not be translated.

## CONTRADICTIONS

RESPONSIBLE `mra`  
STATE `OBSOLET` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-11` ESTIMATEDHOURS `1`  

### Conflict: Rights Management vs responsible

Rights Management is a Non Goal but on the other hand we have
persons responsible for paragraphs. 

If anyone may edit any paragraph, then the responsibility
must constantly change (and is worthless) or it is false.


> `mra`  
>   
> Yes the responsability should change often in order to request feedback.  
> If multiple persons work on a spec, the spec is considered teamwork  
> and the people should work together.  
>   
> Only one person is responsible for each paragraph at a single point in  
> time. This could be enforced by external tools as well as rights management  
> if necessary. Although with version control it should not be necessary.


## UNDECIDED

RESPONSIBLE `mra`  
STATE `OBSOLET` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-11` ESTIMATEDHOURS `1`  

It is not really clear how the linefeed is
exported in markdown and json.
See the linefeed paragraph

> `mra`  
>   
> No longer an issue. Every import and export replaces  
> to `\n` only. Editors should be able to handle it,   
> otherwise other tools would be needed to convert  
> before editing.


## DEFINITIONS

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### All definitions of the `speclib` apply

******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### -name

The parameter passed via `-name` is the name of a spec property (like
`TRANSLATIONS` with the command `property`) or the name of a paragraph
property (like `STATE` with the command `meta`).

******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### -usage

The parameter passed via `-usage` is the name of a filter provided by
`speccer` for typical usage scenarios.

There are currently:

- **documentation** filters 
  - PROPERTIES
  - META
  - COMMENTS
  - NONGOALS
  - CONTRADICTIONS
  - UNDECIDED
  - SPEC_END
  - RESOURCES
  - PLANNING
  - APPROVED
  - PARTLY_IMPLEMENTED
  - OBSOLET
- **approval** filters
  - META
  - COMMENTS
  - CONTRADICTIONS
  - UNDECIDED
  - SPEC_END
  - RESOURCES
  - PLANNING
  - OBSOLET
- **discussion** filters
  - OVERVIEW
  - APPROVED
  - PARTLY_IMPLEMENTED
  - OBSOLET
  - SPEC_END
- **implementation** filters
  - UNDECIDED
  - PLANNING
  - FULLY_IMPLEMENTED
  - OBSOLET


## FEATURES

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### create a new spec

Go into a folder that should contain your spec and run

    speccer -cmd create

Creates a file `spec.json` in the current folder and exports it
as `spec.md` (Markdown) and `spec.html` (HTML).

If the `-language` parameter is set (to e.g. `en_US`) it creates
the following files instead:
- `spec_en_US.json`
- `spec_en_US.md`
- `spec_en_US.html`

Be aware that the new spec is not valid since there are missing
properties. You can s validate a spec to see them:

    speccer -cmd validate



******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### validate a spec

Go into a folder that contains your spec and run

    speccer -cmd validate

It it returns nothing your spec is valid.

******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### recreate HTML and markdown export based on current spec.json

Go into a folder that contains your spec and run

    speccer -cmd save



******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### export markdown to stdout without overwriting json.md

Go into a folder that contains your spec and run

    speccer -cmd markdown

You might also pass the usage option to filter out some data, e.g.

    speccer -cmd markdown -usage "documentation"

    



******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### export HTML to stdout without overwriting json.html

Go into a folder that contains your spec and run

    speccer -cmd html

You might also pass the usage option to filter out some data, e.g.

    speccer -cmd html -usage "documentation"





******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### print property of spec

Go into a folder that contains your spec and run

    speccer -cmd property -name COMPANY

to print the COMPANY property. The properties are 
all written in high caps and defined in `speclib`.




******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### set property of spec

Go into a folder that contains your spec and run

    speccer -cmd property -name COMPANY -set acme

to set the `COMPANY` property to `acme`. The properties are 
all written in high caps and defined in `speclib`.




******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### unset property of spec

Go into a folder that contains your spec and run

    speccer -cmd property -name PARENT -uset

to unset the `PARENT` property. Only the following properties can
be unset:

- REQUESTEDBY
- RELATED
- TRANSLATIONS
- SUPERSEDEDBY
- RESOURCES
- PARENT
- PERSONS


******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### add markdown text to a section

Go into a folder that contains your spec and write your text
into a file, let's say `new.md` and then run

    speccer -cmd add -sec SCENARIOS -resp 'Your Name' -set new.md

to add the content of the file to the SCENARIOS sections.
The available sections are defined in the `speclib`.


******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### show list of paragraphs of a section

Go into a folder that contains your spec and run

    speccer -cmd positions -sec SCENARIOS

to see a list of all paragraphs of the SCENARIOS section.
Each paragraph has the first line printed preceeded by 
a number indicating the position of the paragraph inside
the section. This position can be used to print or change
the paragraph.


******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### show content of a paragraphs of a section

Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to see. Therefor you might want to first 
print all positions of the section (here SCENARIOS) 
with

    speccer -cmd positions -sec SCENARIOS

then you can run

    speccer -cmd text -sec SCENARIOS -at 3

to see the text for the 3rd paragraph in the
SCENARIOS section.

If you want to include comments and meta data, run

    speccer -cmd text -sec SCENARIOS -at 3 -with-comments -with-meta


******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### change content of a paragraphs of a section

Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to see. Therefor you might want to first 
print all positions of the section (here SCENARIOS) 
with

    speccer -cmd positions -sec SCENARIOS

then you can run

    speccer -cmd text -sec SCENARIOS -at 3 > temp.md

to save the current text and comments for the 3rd paragraph in the
SCENARIOS section to the file `temp.md`.

Now fire up your editor to change the content. Then run

    speccer -cmd text -sec SCENARIOS -at 3 -set temp.md

to set the text of the 3rd paragraph in the SCENARIOS section.

The comments and meta data of the 3rd paragraph are not modified 
by this operation. So you may want to change them too.



******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### see, add and change comment to a paragraph of a section

Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to see. Therefor you might want to first 
print all positions of the section (here SCENARIOS) 
with

    speccer -cmd positions -sec SCENARIOS

then you can run

    speccer -cmd comment -sec SCENARIOS -at 3

to see the current comments for the 3rd paragraph in the
SCENARIOS section.

Now if the author of your comment already made a comment to 
this paragraph and if you want to append to this old comment,
you first have to save the old one.

To do this, run 

    speccer -cmd comment -sec SCENARIOS -at 3 -author Jim > temp.md

to save the comment of `Jim` to the file `temp.md`. Then you can edit this
file. If you just want to remove the old comment or if there was no
comment you may simply start with an empty file `temp.md`.

After you edited the file, run 

    speccer -cmd comment -sec SCENARIOS -at 3 -author Jim -set temp.md

to set the comment of `Jim` on the 3rd paragraph of the SCENARIOS section
to the content of the file `temp.md`.


******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### move a paragraph of a section

Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to move. Therefor you might want to first 
print all positions of the section (here SCENARIOS) 
with

    speccer -cmd positions -sec SCENARIOS

then you can run

    speccer -cmd move -sec SCENARIOS -at 3 -to 1

to move the 3rd paragraph in the SCENARIOS section to be the first.


******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### delete a paragraph of a section

Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to delete. Therefor you might want to first 
print all positions of the section (here SCENARIOS) 
with

    speccer -cmd positions -sec SCENARIOS

then you can run

    speccer -cmd rm -sec SCENARIOS -at 3

to remove the 3rd paragraph in the SCENARIOS section.


******

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

### see and change meta data of a paragraph of a section

Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to see the meta data. Therefor you might want to first 
print all positions of the section (here SCENARIOS) 
with

    speccer -cmd positions -sec SCENARIOS

then you can run

    speccer -cmd meta -sec SCENARIOS -at 3

to see all meta data of the 3rd paragraph in the SCENARIOS section.

If you want to change a datum you will have to pass the name:

    speccer -cmd meta -sec SCENARIOS -at 3 -name RESPONSIBLE -set Jim




# SPEC_END
