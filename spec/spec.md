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




## SCENARIO 1 Create Specification

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Mike wants to create a specification. He gets a clean
starting point with `speccer`.

******


## SCENARIO 2 Add comments

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Mike receives comments for his specification.
He wants to add them without hassle.


******


## SCENARIO 3 See and update specification

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Mike wants to update his specification with is
editor.


******


## SCENARIO 4 Script

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Gill wants to change specifications with a script.


******


## SCENARIO 5 Import and export

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Gill wants to read import and export specifications
from different places.


******


## SCENARIO 6 View as HTML and markdown

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Sue wants to read the specifications in HTML and 
be able to mail them as markdown.


******


## SCENARIO 7 Filter for documentation

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Peter wants to create a documentation from the spec and
wants the unneccessary informations to be filtered.

******


## SCENARIO 8 Filter for implementation

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Anne wants to implement parts of the specification.
She just wants to read the parts that are ready
to be implemented.

## NONGOAL 1 No GUI

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

To have a GUI would be a project on its own.
However such a project could also be external and be built upon speccer.

******


## NONGOAL 2 No Database

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

`speccer` will operate on plain text json files only.
There will be no way to have a database and search.
A database could however exchange data with `speccer`
via json files.

******


## NONGOAL 3 No versioning

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

`speccer` will not do any backups or have versions.
You are advised to use a version control system to save your
specs or make backups by a backup tool.

******


## NONGOAL 4 No editing

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

`speccer` will have any functionality of an editor.
You may however replace parts of the spec by exporting them
to your favourite editor and reimport them when you are finished.

******


## NONGOAL 5 No merging of specs

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

`speccer` won't provide any mechanism to merge specs.
Since you can export and import parts of the specs in
plain text or markdown, it is not hard to merge by hand.

******


## NONGOAL 6 No copying of specs

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

`speccer` won't help you copying specs, since you can simply
copy the json files yourself and then edit the copy.

******


## NONGOAL 7 No rights management

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

`speccer` won't have any rights management baked in. 
You could manage the viewing and changing rights for
a single spec by the filesystem.

If you need finer control, you need to write a wrapper
around `speccer` or the `speclib`.

******


## NONGOAL 8 No multiple comments per author and paragraph

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

With `speclib` (and therefor `speccer` too),
a single author can only have one comment per paragraph.

We consider this a good thing, since it forces the spec
to be split up in several paragraphs if they get too many
comments. The spec will benefit from it.

Also it is possible to have comments of arbitrary length.


******


## NONGOAL 9 No 'threads'

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Again this is a decision by `speclib` and considered a good 
thing. There should not be comments on comments and endless
discussions. Instead the comments should be directed to a 
paragraph and decision should be made soon.

Then the different topics should seperate into the different
sections like 'UNDECIDED', 'CONTRADICTIONS', 'DEFINITIONS',
'FEATURES' etc. and there should be multiple paragraphs 
which may each have their own state of consense and comments.


******


## NONGOAL 10 No order of comments

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Again this is a decision by `speclib`.
Since each author may just have one comment per paragraph, 
there is no point in have an order for comments.

However their is an order of paragraphs and that could 
be changed.


******


## NONGOAL 11 No text generation

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Since the specs format is based on json, it is very easy 
and flexible to have own templates and copy them.

******


## NONGOAL 12 No PDF export

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

There are great tools to convert from HTML to PDF, like `wkhtmltopdf` and
also markdown to pdf converter like `pandoc`. Since `speccer` exports 
to HTML and markdown you should be able to use these tools to generate
PDFs.

******


## NONGOAL 13 No Translations for SPECS

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

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


## NONGOAL 14 No time tracking

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Each paragraph has a field ESTIMATEDHOURS that could be used to
do basic scheduling.

An external tool could then collect the estimated hours and compare
them to the real time to drive conclusions about the schedule.

More time tracking functionality is not part of this project.


******


## NONGOAL 15 No shortcuts for commands or presets

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Some commands and options are hard to remember.
But instead of bake new options into `speccer` we
deliver an alias file you could use in your bash or zsh shell.

Have a look at it to see how easy you could do one on your own.
In the alias you could also set defaults for responsibles etc.


******


## NONGOAL 16 No translation of options

RESPONSIBLE `mra`  
STATE `APPROVED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

In order to be able to associate the wordings of the options
to the corresponding keys in the json file and the untranslated
specs the options will not be translated.

## CONTRADICTION 1 Rights Management vs responsible

RESPONSIBLE `mra`  
STATE `OBSOLET` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-11` ESTIMATEDHOURS `1`  

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


## UNDECIDED 1 It is not really clear how the linefeed is

RESPONSIBLE `mra`  
STATE `OBSOLET` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-11` ESTIMATEDHOURS `1`  

exported in markdown and json.
See the linefeed paragraph

> `mra`  
>   
> No longer an issue. Every import and export replaces  
> to `\n` only. Editors should be able to handle it,   
> otherwise other tools would be needed to convert  
> before editing.


## DEFINITION 1 All definitions of the `speclib` apply

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  



******


## DEFINITION 2 -name

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

The parameter passed via `-name` is the name of a spec property (like
`TRANSLATIONS` with the command `property`) or the name of a paragraph
property (like `STATE` with the command `meta`).

******


## DEFINITION 3 -usage

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

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


******


## DEFINITION 4 Current Spec

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-13` ESTIMATEDHOURS `1`  


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

## FEATURE 1 create a new spec

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

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
properties. You can validate a spec to see them:

    speccer -cmd validate





******


## FEATURE 2 validate a spec

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec and run

    speccer -cmd validate

It it returns nothing your spec is valid.

******


## FEATURE 3 recreate HTML and markdown export based on current spec.json

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec and run

    speccer -cmd save



******


## FEATURE 4 export markdown to stdout without overwriting json.md

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec and run

    speccer -cmd markdown

You might also pass the usage option to filter out some data, e.g.

    speccer -cmd markdown -usage "documentation"

    



******


## FEATURE 5 export HTML to stdout without overwriting json.html

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec and run

    speccer -cmd html

You might also pass the usage option to filter out some data, e.g.

    speccer -cmd html -usage "documentation"





******


## FEATURE 6 print property of spec

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec and run

    speccer -cmd property -name COMPANY

to print the COMPANY property. The properties are 
all written in high caps and defined in `speclib`.




******


## FEATURE 7 set property of spec

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec and run

    speccer -cmd property -name COMPANY -set acme

to set the `COMPANY` property to `acme`. The properties are 
all written in high caps and defined in `speclib`.




******


## FEATURE 8 unset property of spec

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-12`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

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


## FEATURE 9 add markdown text to a section

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec and write your text
into a file, let's say `new.md`.

Put the title of the paragraph in the first line 
and the body in the next lines.
Then run

    speccer -cmd add -sec SCENARIO -resp 'Your Name' -set new.md

to add title and body to the SCENARIO sections.
The available sections are defined in the `speclib`.






******


## FEATURE 10 show list of paragraphs of a section

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec and run

    speccer -cmd positions -sec SCENARIO

to see a list of all paragraphs of the SCENARIO section.
Each paragraph has the first line printed preceeded by 
a number indicating the position of the paragraph inside
the section. This position can be used to print or change
the paragraph.




******


## FEATURE 11 show content of a paragraphs of a section

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to see. Therefor you might want to first 
print all positions of the section (here SCENARIO) 
with

    speccer -cmd positions -sec SCENARIO

then you can run

    speccer -cmd text -sec SCENARIO -at 3

to see the text for the 3rd paragraph in the
SCENARIO section.

If you want to include comments and meta data, run

    speccer -cmd text -sec SCENARIO -at 3 -with-comments -with-meta




******


## FEATURE 12 change content of a paragraph of a section

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to see. Therefor you might want to first 
print all positions of the section (here SCENARIO) 
with

    speccer -cmd positions -sec SCENARIO

then you can run

    speccer -cmd text -sec SCENARIO -at 3 > temp.md

to save the current text and comments for the 3rd paragraph in the
SCENARIO section to the file `temp.md`.

Now fire up your editor to change the content. Then run

    speccer -cmd text -sec SCENARIO -at 3 -set temp.md

to set the text of the 3rd paragraph in the SCENARIO section.

The comments and meta data (including the title) of the 3rd paragraph are not modified by this operation. So you may want to change them too.







******


## FEATURE 13 see, add and change comment to a paragraph of a section

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to see. Therefor you might want to first 
print all positions of the section (here SCENARIO) 
with

    speccer -cmd positions -sec SCENARIO

then you can run

    speccer -cmd comment -sec SCENARIO -at 3

to see the current comments for the 3rd paragraph in the
SCENARIO section.

Now if the author of your comment already made a comment to 
this paragraph and if you want to append to this old comment,
you first have to save the old one.

To do this, run 

    speccer -cmd comment -sec SCENARIO -at 3 -author Jim > temp.md

to save the comment of `Jim` to the file `temp.md`. Then you can edit this
file. If you just want to remove the old comment or if there was no
comment you may simply start with an empty file `temp.md`.

After you edited the file, run 

    speccer -cmd comment -sec SCENARIO -at 3 -author Jim -set temp.md

to set the comment of `Jim` on the 3rd paragraph of the SCENARIO section
to the content of the file `temp.md`.




******


## FEATURE 14 move a paragraph of a section

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to move. Therefor you might want to first 
print all positions of the section (here SCENARIO) 
with

    speccer -cmd positions -sec SCENARIO

then you can run

    speccer -cmd move -sec SCENARIO -at 3 -to 1

to move the 3rd paragraph in the SCENARIO section to be the first.




******


## FEATURE 15 delete a paragraph of a section

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to delete. Therefor you might want to first 
print all positions of the section (here SCENARIO) 
with

    speccer -cmd positions -sec SCENARIO

then you can run

    speccer -cmd rm -sec SCENARIO -at 3

to remove the 3rd paragraph in the SCENARIO section.




******


## FEATURE 16 see and change meta data of a paragraph of a section

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-12` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec.
You will need to know the position of the paragraph
you want to see the meta data. Therefor you might want to first 
print all positions of the section (here SCENARIO) 
with

    speccer -cmd positions -sec SCENARIO

then you can run

    speccer -cmd meta -sec SCENARIO -at 3

to see all meta data of the 3rd paragraph in the SCENARIO section.

If you want to change a datum you will have to pass the name:

    speccer -cmd meta -sec SCENARIO -at 3 -name RESPONSIBLE -set Jim






******


## FEATURE 17 shortcuts

RESPONSIBLE `mra`  
STATE `FULLY_IMPLEMENTED` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-13` ESTIMATEDHOURS `1`  

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
`spec_save`   

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
`spec_add-SCENARIO`  **new.md**  
`spec_add-CONTRADICTION`  **new.md**  
`spec_add-DEFINITION` **new.md**  
`spec_add-FEATURE` **new.md**  
`spec_add-NONGOAL`  **new.md**  
`spec_add-UNDECIDED` **new.md**  

#### Show Sections
`spec-OVERVIEW`   
`spec-SCENARIO-at` **2**  
`spec-CONTRADICTION-at` **5**  
`spec-DEFINITION-at` **20**  
`spec-FEATURE-at` **1**  
`spec-NONGOAL-at`  **2**  
`spec-UNDECIDED-at`  **5**  

#### Show Sections with comments and meta
`spec-OVERVIEW-full`  
`spec-SCENARIO-full-at` **2**  
`spec-CONTRADICTION-full-at` **2**  
`spec-DEFINITION-full-at` **5**  
`spec-FEATURE-full-at` **2**  
`spec-NONGOAL-full-at`  **7**  
`spec-UNDECIDED-full-at` **2**  

#### Get list of Paragraphs for all section
`spec_ls`  

#### Get list of Paragraphs for a section
`spec_ls-SCENARIO`   
`spec_ls-CONTRADICTION`   
`spec_ls-DEFINITION`   
`spec_ls-FEATURE`   
`spec_ls-NONGOAL`   
`spec_ls-UNDECIDED`   

#### Get Comments of Paragraphs for a section
`spec_comment-OVERVIEW`   
`spec_comment-SCENARIO-at`  **2**  
`spec_comment-CONTRADICTION-at`  **2**  
`spec_comment-DEFINITION-at`  **2**  
`spec_comment-FEATURE-at`  **2**  
`spec_comment-NONGOAL-at`  **2**  
`spec_comment-UNDECIDED-at` **2**  

#### Set Comment of Paragraph for a section
`spec_comment_replace-OVERVIEW`  **new.md**  
`spec_comment_replace-SCENARIO`  **new.md -at 2**  
`spec_comment_replace-CONTRADICTION`  **new.md -at 2**  
`spec_comment_replace-DEFINITION`  **new.md -at 2**  
`spec_comment_replace-FEATURE`  **new.md -at 2**  
`spec_comment_replace-NONGOAL`  **new.md -at 2**  
`spec_comment_replace-UNDECIDED`  **new.md -at 2**  

#### Replace Sections
`spec_replace-OVERVIEW`  **new.md**  
`spec_replace-SCENARIO`  **new.md -at 2**  
`spec_replace-CONTRADICTION` **new.md -at 2**  
`spec_replace-DEFINITION` **new.md -at 2**  
`spec_replace-FEATURE` **new.md -at 2**  
`spec_replace-NONGOAL` **new.md -at 2**  
`spec_replace-UNDECIDED` **new.md -at 2**  

#### Move Sections
`spec_mv-SCENARIO-at`  **2 -to 5**  
`spec_mv-CONTRADICTION-at` **2 -to 5**  
`spec_mv-DEFINITION-at` **2 -to 5**  
`spec_mv-FEATURE-at` **2 -to 5**  
`spec_mv-NONGOAL-at` **2 -to 5**  
`spec_mv-UNDECIDED-at` **2 -to 5**  

#### Remove Sections
`spec_rm-SCENARIO-at` **2**  
`spec_rm-CONTRADICTION-at` **2**  
`spec_rm-DEFINITION-at` **2**  
`spec_rm-FEATURE-at` **2**  
`spec_rm-NONGOAL-at` **2**  
`spec_rm-UNDECIDED-at`  **2**  

#### Get Meta Data of Overview / Spec
`spec-RESPONSIBLE`  
`spec-STATE`  
`spec-DEADLINE`  
`spec-LASTUPDATE`  
`spec-ESTIMATEDHOURS`  
`spec-TITLE`  

#### Get Meta Data of Section
`spec-SCENARIO-RESPONSIBLE-at`  **2**  
`spec-SCENARIO-STATE-at`  **2**  
`spec-SCENARIO-DEADLINE-at`  **2**  
`spec-SCENARIO-LASTUPDATE-at`  **2**  
`spec-SCENARIO-ESTIMATEDHOURS-at`  **2**  
`spec-SCENARIO-TITLE-at`  **2**  

`spec-CONTRADICTION-RESPONSIBLE-at`  **2**  
`spec-CONTRADICTION-STATE-at`  **2**  
`spec-CONTRADICTION-DEADLINE-at`  **2**  
`spec-CONTRADICTION-LASTUPDATE-at`  **2**  
`spec-CONTRADICTION-ESTIMATEDHOURS-at`  **2**  
`spec-CONTRADICTION-TITLE-at`  **2**  

`spec-DEFINITION-RESPONSIBLE-at`  **2**  
`spec-DEFINITION-STATE-at`  **2**  
`spec-DEFINITION-DEADLINE-at`  **2**  
`spec-DEFINITION-LASTUPDATE-at`  **2**  
`spec-DEFINITION-ESTIMATEDHOURS-at`  **2**  
`spec-DEFINITION-TITLE-at`  **2**  

`spec-FEATURE-RESPONSIBLE-at`  **2**  
`spec-FEATURE-STATE-at`  **2**  
`spec-FEATURE-DEADLINE-at`  **2**  
`spec-FEATURE-LASTUPDATE-at`  **2**  
`spec-FEATURE-ESTIMATEDHOURS-at`  **2**  
`spec-FEATURE-TITLE-at`  **2**  

`spec-NONGOAL-RESPONSIBLE-at`  **2**  
`spec-NONGOAL-STATE-at`  **2**  
`spec-NONGOAL-DEADLINE-at`   **2**  
`spec-NONGOAL-LASTUPDATE-at`  **2**  
`spec-NONGOAL-ESTIMATEDHOURS-at`  **2**  
`spec-NONGOAL-TITLE-at`  **2**  

`spec-UNDECIDED-RESPONSIBLE-at`  **2**  
`spec-UNDECIDED-STATE-at`  **2**  
`spec-UNDECIDED-DEADLINE-at`  **2**  
`spec-UNDECIDED-LASTUPDATE-at`  **2**  
`spec-UNDECIDED-ESTIMATEDHOURS-at` **2**  
`spec-UNDECIDED-TITLE-at`  **2**  

#### Set Meta Data of Overview / Spec
`spec_set-RESPONSIBLE` **'Peter Pan'**  
`spec_set-PLANNING`  
`spec_set-APPROVED`  
`spec_set-PARTLY_IMPLEMENTED`  
`spec_set-FULLY_IMPLEMENTED`   
`spec_set-OBSOLET`  
`spec_set-DEADLINE`  **'2014-12-24'**  
`spec_set-ESTIMATEDHOURS`  **12**  
`spec_set-TITLE`  **'my title'**  

#### Set Meta Data of Section

`spec_set-SCENARIO-RESPONSIBLE`   **'Peter Pan' -at 2**    
`spec_set-SCENARIO-PLANNING`  **-at 2**  
`spec_set-SCENARIO-APPROVED`   **-at 2**  
`spec_set-SCENARIO-PARTLY_IMPLEMENTED`   **-at 2**  
`spec_set-SCENARIO-FULLY_IMPLEMENTED`   **-at 2**  
`spec_set-SCENARIO-OBSOLET`   **-at 2**  
`spec_set-SCENARIO-DEADLINE`   **'2014-12-24' -at 2**  
`spec_set-SCENARIO-ESTIMATEDHOURS`  **12 -at 2**  
`spec_set-SCENARIO-TITLE`  **'my title' -at 3**  

`spec_set-CONTRADICTION-RESPONSIBLE`  **'Peter Pan' -at 2**    
`spec_set-CONTRADICTION-PLANNING`   **-at 2**  
`spec_set-CONTRADICTION-APPROVED`   **-at 2**  
`spec_set-CONTRADICTION-PARTLY_IMPLEMENTED`   **-at 2**  
`spec_set-CONTRADICTION-FULLY_IMPLEMENTED`   **-at 2**  
`spec_set-CONTRADICTION-OBSOLET`   **-at 2**  
`spec_set-CONTRADICTION-DEADLINE`  **'2014-12-24' -at 2**  
`spec_set-CONTRADICTION-ESTIMATEDHOURS`  **3 -at 2**  
`spec_set-CONTRADICTION-TITLE`  **'my title' -at 3**  

`spec_set-DEFINITION-RESPONSIBLE` **'Peter Pan' -at 2**    
`spec_set-DEFINITION-PLANNING`   **-at 2**  
`spec_set-DEFINITION-APPROVED`   **-at 2**  
`spec_set-DEFINITION-PARTLY_IMPLEMENTED`   **-at 2**  
`spec_set-DEFINITION-FULLY_IMPLEMENTED`   **-at 2**  
`spec_set-DEFINITION-OBSOLET`   **-at 2**  
`spec_set-DEFINITION-DEADLINE`  **'2014-12-24' -at 2**  
`spec_set-DEFINITION-ESTIMATEDHOURS`  **3 -at 2**  
`spec_set-DEFINITION-TITLE`  **'my title' -at 3**  

`spec_set-FEATURE-RESPONSIBLE`  **'Peter Pan' -at 2**    
`spec_set-FEATURE-PLANNING`  **-at 2**  
`spec_set-FEATURE-APPROVED`  **-at 2**  
`spec_set-FEATURE-PARTLY_IMPLEMENTED`  **-at 2**  
`spec_set-FEATURE-FULLY_IMPLEMENTED`  **-at 2**  
`spec_set-FEATURE-OBSOLET`   **-at 2**  
`spec_set-FEATURE-DEADLINE`  **'2014-12-24' -at 2**  
`spec_set-FEATURE-ESTIMATEDHOURS`  **3 -at 2**  
`spec_set-FEATURE-TITLE`  **'my title' -at 3**  

`spec_set-NONGOAL-RESPONSIBLE` **'Peter Pan' -at 2**    
`spec_set-NONGOAL-PLANNING`   **-at 2**  
`spec_set-NONGOAL-APPROVED`   **-at 2**  
`spec_set-NONGOAL-PARTLY_IMPLEMENTED`  **-at 2**  
`spec_set-NONGOAL-FULLY_IMPLEMENTED`   **-at 2**  
`spec_set-NONGOAL-OBSOLET`   **-at 2**  
`spec_set-NONGOAL-DEADLINE`  **'2014-12-24' -at 2**  
`spec_set-NONGOAL-ESTIMATEDHOURS`  **3 -at 2**  
`spec_set-NONGOAL-TITLE`  **'my title' -at 3**  

`spec_set-UNDECIDED-RESPONSIBLE`  **'Peter Pan' -at 2**    
`spec_set-UNDECIDED-PLANNING`   **-at 2**  
`spec_set-UNDECIDED-APPROVED`   **-at 2**  
`spec_set-UNDECIDED-PARTLY_IMPLEMENTED`   **-at 2**  
`spec_set-UNDECIDED-FULLY_IMPLEMENTED`   **-at 2**  
`spec_set-UNDECIDED-OBSOLET`   **-at 2**  
`spec_set-UNDECIDED-DEADLINE`  **'2014-12-24' -at 2**  
`spec_set-UNDECIDED-ESTIMATEDHOURS`  **3 -at 2**  
`spec_set-UNDECIDED-TITLE`  **'my title' -at 3**  



******


## FEATURE 18 adding multiple paragraphs to a section in one go

RESPONSIBLE `benny`  
STATE `PLANNING` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-13` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec.

Create a new file (e.g. `new.md`) and write your paragraphs
into this file.

Each paragraph should hold its title inside its first line
and should end with a line containing `***`.

Then run

    speccer -cmd add-multi -sec SCENARIO -resp Jim -set new.md

to add multiple paragraphs to the SCENARIO section as defined in
`new.md` with all having the responsible set to Jim.



******


## FEATURE 19 setting multiple comments to a paragraph in one go

RESPONSIBLE `benny`  
STATE `PLANNING` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-13` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec.

Create a new file (e.g. `new.md`) and write the comments
into this file.

Each comment should hold its author inside the first line
and should end with a line containing `***`

You will need to know the position of the paragraph
you want to set the comments to. Therefor you might want to first 
print all positions of the section (here SCENARIO) 
with

    speccer -cmd positions -sec SCENARIO

then you can run

    speccer -cmd comment-multi -sec SCENARIO -at 3 -set new.md

to set multiple comments to the 3rd paragraph of the SCENARIO section as defined in `new.md`.

    

******


## FEATURE 20 packing multiple paragraphs to one single paragraph

RESPONSIBLE `benny`  
STATE `PLANNING` LASTUPDATE `2013-11-13`  
DEADLINE `2013-11-13` ESTIMATEDHOURS `1`  

Go into a folder that contains your spec.

You will need to know the positions of the paragraph you want
to pack. Therefor you might want to first 
print all positions of the section (here SCENARIO) 
with

    speccer -cmd positions -sec SCENARIO

then you can run

    speccer -cmd pack -sec SCENARIO -packing '[3,4,5]' -resp Jim -title 'all together'

to pack the paragraphs no 3,4 and 5 of the SCENARIO section into a
new paragraph with the title 'all together' and the responsible 'Jim'.
The previous titles are inserted as subheadings and the previous paragraphs
are deleted.


# SPEC_END
