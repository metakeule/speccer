# Common
alias spec-create='speccer -cmd create';
alias spec-validate='speccer -cmd validate';
alias spec-save='speccer -cmd save';

alias spec-markdown='speccer -cmd markdown';
alias spec-markdown-doc='speccer -cmd markdown -usage "documentation"';
alias spec-markdown-approve='speccer -cmd markdown -usage "approval"';
alias spec-markdown-discuss='speccer -cmd markdown -usage "discussion"';
alias spec-markdown-implement='speccer -cmd markdown -usage "implementation"';
alias spec-markdown-clean='speccer -cmd markdown -filter '"'"'{"META": true, "COMMENTS": true}'"'"'';

alias spec-html='speccer -cmd html';
alias spec-html-doc='speccer -cmd html -usage "documentation"';
alias spec-html-approve='speccer -cmd html -usage "approval"';
alias spec-html-discuss='speccer -cmd html -usage "discussion"';
alias spec-html-implement='speccer -cmd html -usage "implementation"';
alias spec-html-clean='speccer -cmd html -filter '"'"'{"META": true, "COMMENTS": true}'"'"'';

# Spec property getters
alias spec-REQUESTEDBY='speccer -cmd property -name REQUESTEDBY';  
alias spec-RELATED='speccer -cmd property -name RELATED';
alias spec-TRANSLATIONS='speccer -cmd property -name TRANSLATIONS'; 
alias spec-SUPERSEDEDBY='speccer -cmd property -name SUPERSEDEDBY'; 
alias spec-RESOURCES='speccer -cmd property -name RESOURCES';
alias spec-PARENT='speccer -cmd property -name PARENT';
alias spec-COMPANY='speccer -cmd property -name COMPANY';
alias spec-PROJECT='speccer -cmd property -name PROJECT';
alias spec-PERSONS='speccer -cmd property -name PERSONS';
alias spec-URL='speccer -cmd property -name URL';
alias spec-LANGUAGE='speccer -cmd property -name LANGUAGE';
alias spec-DATEFORMAT='speccer -cmd property -name DATEFORMAT';
alias spec-APPROVED='speccer -cmd property -name APPROVED';     

# Spec property unsetters
alias spec-REQUESTEDBY-uset='speccer -cmd property -name REQUESTEDBY -uset';  
alias spec-RELATED-uset='speccer -cmd property -name RELATED -uset';
alias spec-TRANSLATIONS-uset='speccer -cmd property -name TRANSLATIONS -uset'; 
alias spec-SUPERSEDEDBY-uset='speccer -cmd property -name SUPERSEDEDBY -uset'; 
alias spec-RESOURCES-uset='speccer -cmd property -name RESOURCES -uset';
alias spec-PARENT-uset='speccer -cmd property -name PARENT -uset';
alias spec-PERSONS-uset='speccer -cmd property -name PERSONS -uset';

# Spec property setters
alias spec-REQUESTEDBY-set='speccer -cmd property -name REQUESTEDBY -set';  
alias spec-RELATED-set='speccer -cmd property -name RELATED -set';
alias spec-TRANSLATIONS-set='speccer -cmd property -name TRANSLATIONS -set'; 
alias spec-SUPERSEDEDBY-set='speccer -cmd property -name SUPERSEDEDBY -set'; 
alias spec-RESOURCES-set='speccer -cmd property -name RESOURCES -set';
alias spec-PARENT-set='speccer -cmd property -name PARENT -set';
alias spec-COMPANY-set='speccer -cmd property -name COMPANY -set';
alias spec-PROJECT-set='speccer -cmd property -name PROJECT -set';

alias spec-URL-set='speccer -cmd property -name URL -set';
alias spec-LANGUAGE-set='speccer -cmd property -name LANGUAGE -set';
alias spec-DATEFORMAT-set='speccer -cmd property -name DATEFORMAT -set';
alias spec-APPROVED-set='speccer -cmd property -name APPROVED -set';     
alias spec-PERSONS-set='speccer -cmd property -name PERSONS -set';     

# Add Sections
alias spec-add-SCENARIO="speccer -cmd add -sec SCENARIO -resp $USER -set";
alias spec-add-CONTRADICTION="speccer -cmd add -sec CONTRADICTION -resp $USER -set";
alias spec-add-DEFINITION="speccer -cmd add -sec DEFINITION -resp $USER -set";
alias spec-add-FEATURE="speccer -cmd add -sec FEATURE -resp $USER -set";
alias spec-add-NONGOAL="speccer -cmd add -sec NONGOAL -resp $USER -set";
alias spec-add-UNDECIDED="speccer -cmd add -sec UNDECIDED -resp $USER -set";

# Show Sections
alias spec-OVERVIEW='speccer -cmd text -sec OVERVIEW';
alias spec-SCENARIO-at='speccer -cmd text -sec SCENARIO -at';
alias spec-CONTRADICTION-at='speccer -cmd text -sec CONTRADICTION -at';
alias spec-DEFINITION-at='speccer -cmd text -sec DEFINITION -at';
alias spec-FEATURE-at='speccer -cmd text -sec FEATURE -at';
alias spec-NONGOAL-at='speccer -cmd text -sec NONGOAL -at';
alias spec-UNDECIDED-at='speccer -cmd text -sec UNDECIDED -at';

alias spec-OVERVIEW-full='speccer -cmd text -sec OVERVIEW -with-comments -with-meta';
alias spec-SCENARIO-full-at='speccer -cmd text -sec SCENARIO -with-comments -with-meta -at';
alias spec-CONTRADICTION-full-at='speccer -cmd text -sec CONTRADICTION -with-comments -with-meta -at';
alias spec-DEFINITION-full-at='speccer -cmd text -sec DEFINITION -with-comments -with-meta -at';
alias spec-FEATURE-full-at='speccer -cmd text -sec FEATURE -with-comments -with-meta -at';
alias spec-NONGOAL-full-at='speccer -cmd text -sec NONGOAL -with-comments -with-meta -at';
alias spec-UNDECIDED-full-at='speccer -cmd text -sec UNDECIDED -with-comments -with-meta -at';

# Get list of Paragraphs for a section
alias spec-ls='speccer -cmd positions';
alias spec-ls-SCENARIO='speccer -cmd positions -sec SCENARIO';
alias spec-ls-CONTRADICTION='speccer -cmd positions -sec CONTRADICTION';
alias spec-ls-DEFINITION='speccer -cmd positions -sec DEFINITION';
alias spec-ls-FEATURE='speccer -cmd positions -sec FEATURE';
alias spec-ls-NONGOAL='speccer -cmd positions -sec NONGOAL';
alias spec-ls-UNDECIDED='speccer -cmd positions -sec UNDECIDED';

# Get Comments of Paragraphs for a section
alias spec-comment-OVERVIEW='speccer -cmd comment -sec OVERVIEW';
alias spec-comment-SCENARIO-at='speccer -cmd comment -sec SCENARIO -at';
alias spec-comment-CONTRADICTION-at='speccer -cmd comment -sec CONTRADICTION -at';
alias spec-comment-DEFINITION-at='speccer -cmd comment -sec DEFINITION -at';
alias spec-comment-FEATURE-at='speccer -cmd comment -sec FEATURE -at';
alias spec-comment-NONGOAL-at='speccer -cmd comment -sec NONGOAL -at';
alias spec-comment-UNDECIDED-at='speccer -cmd comment -sec UNDECIDED -at';

# Set Comment of Paragraph for a section
alias spec-comment-OVERVIEW-set='speccer -cmd comment -sec OVERVIEW -set';
alias spec-comment-SCENARIO-set='speccer -cmd comment -sec SCENARIO -set';
alias spec-comment-CONTRADICTION-set='speccer -cmd comment -sec CONTRADICTION -set';
alias spec-comment-DEFINITION-set='speccer -cmd comment -sec DEFINITION -set';
alias spec-comment-FEATURE-set='speccer -cmd comment -sec FEATURE -set';
alias spec-comment-NONGOAL-set='speccer -cmd comment -sec NONGOAL -set';
alias spec-comment-UNDECIDED-set='speccer -cmd comment -sec UNDECIDED -set';

# Replace Sections
alias spec-set-OVERVIEW='speccer -cmd text -sec OVERVIEW -set';
alias spec-set-SCENARIO='speccer -cmd text -sec SCENARIO -set';
alias spec-set-CONTRADICTION='speccer -cmd text -sec CONTRADICTION -set';
alias spec-set-DEFINITION='speccer -cmd text -sec DEFINITION -set';
alias spec-set-FEATURE='speccer -cmd text -sec FEATURE -set';
alias spec-set-NONGOAL='speccer -cmd text -sec NONGOAL -set';
alias spec-set-UNDECIDED='speccer -cmd text -sec UNDECIDED -set';

# Move Sections
alias spec-mv-SCENARIO-at='speccer -cmd move -sec SCENARIO -at';
alias spec-mv-CONTRADICTION-at='speccer -cmd move -sec CONTRADICTION -at';
alias spec-mv-DEFINITION-at='speccer -cmd move -sec DEFINITION -at';
alias spec-mv-FEATURE-at='speccer -cmd move -sec FEATURE -at';
alias spec-mv-NONGOAL-at='speccer -cmd move -sec NONGOAL -at';
alias spec-mv-UNDECIDED-at='speccer -cmd move -sec UNDECIDED -at';

# Remove Sections
alias spec-rm-SCENARIO-at='speccer -cmd rm -sec SCENARIO -at';
alias spec-rm-CONTRADICTION-at='speccer -cmd rm -sec CONTRADICTION -at';
alias spec-rm-DEFINITION-at='speccer -cmd rm -sec DEFINITION -at';
alias spec-rm-FEATURE-at='speccer -cmd rm -sec FEATURE -at';
alias spec-rm-NONGOAL-at='speccer -cmd rm -sec NONGOAL -at';
alias spec-rm-UNDECIDED-at='speccer -cmd rm -sec UNDECIDED -at';

# Get Meta Data from Section
alias spec-RESPONSIBLE='speccer -cmd meta -sec OVERVIEW -name RESPONSIBLE';
alias spec-STATE='speccer -cmd meta -sec OVERVIEW -name STATE';
alias spec-DEADLINE='speccer -cmd meta -sec OVERVIEW -name DEADLINE';
alias spec-LASTUPDATE='speccer -cmd meta -sec OVERVIEW -name LASTUPDATE';
alias spec-ESTIMATEDHOURS='speccer -cmd meta -sec OVERVIEW -name ESTIMATEDHOURS';
alias spec-TITLE='speccer -cmd meta -sec OVERVIEW -name TITLE';

alias spec-SCENARIO-RESPONSIBLE-at='speccer -cmd meta -sec SCENARIO -name RESPONSIBLE -at';
alias spec-SCENARIO-STATE-at='speccer -cmd meta -sec SCENARIO -name STATE -at';
alias spec-SCENARIO-DEADLINE-at='speccer -cmd meta -sec SCENARIO -name DEADLINE -at';
alias spec-SCENARIO-LASTUPDATE-at='speccer -cmd meta -sec SCENARIO -name LASTUPDATE -at';
alias spec-SCENARIO-ESTIMATEDHOURS-at='speccer -cmd meta -sec SCENARIO -name ESTIMATEDHOURS -at';
alias spec-SCENARIO-TITLE-at='speccer -cmd meta -sec SCENARIO -name TITLE -at';

alias spec-CONTRADICTION-RESPONSIBLE-at='speccer -cmd meta -sec CONTRADICTION -name RESPONSIBLE -at';
alias spec-CONTRADICTION-STATE-at='speccer -cmd meta -sec CONTRADICTION -name STATE -at';
alias spec-CONTRADICTION-DEADLINE-at='speccer -cmd meta -sec CONTRADICTION -name DEADLINE -at';
alias spec-CONTRADICTION-LASTUPDATE-at='speccer -cmd meta -sec CONTRADICTION -name LASTUPDATE -at';
alias spec-CONTRADICTION-ESTIMATEDHOURS-at='speccer -cmd meta -sec CONTRADICTION -name ESTIMATEDHOURS -at';
alias spec-CONTRADICTION-TITLE-at='speccer -cmd meta -sec CONTRADICTION -name TITLE -at';

alias spec-DEFINITION-RESPONSIBLE-at='speccer -cmd meta -sec DEFINITION -name RESPONSIBLE -at';
alias spec-DEFINITION-STATE-at='speccer -cmd meta -sec DEFINITION -name STATE -at';
alias spec-DEFINITION-DEADLINE-at='speccer -cmd meta -sec DEFINITION -name DEADLINE -at';
alias spec-DEFINITION-LASTUPDATE-at='speccer -cmd meta -sec DEFINITION -name LASTUPDATE -at';
alias spec-DEFINITION-ESTIMATEDHOURS-at='speccer -cmd meta -sec DEFINITION -name ESTIMATEDHOURS -at';
alias spec-DEFINITION-TITLE-at='speccer -cmd meta -sec DEFINITION -name TITLE -at';

alias spec-FEATURE-RESPONSIBLE-at='speccer -cmd meta -sec FEATURE -name RESPONSIBLE -at';
alias spec-FEATURE-STATE-at='speccer -cmd meta -sec FEATURE -name STATE -at';
alias spec-FEATURE-DEADLINE-at='speccer -cmd meta -sec FEATURE -name DEADLINE -at';
alias spec-FEATURE-LASTUPDATE-at='speccer -cmd meta -sec FEATURE -name LASTUPDATE -at';
alias spec-FEATURE-ESTIMATEDHOURS-at='speccer -cmd meta -sec FEATURE -name ESTIMATEDHOURS -at';
alias spec-FEATURE-TITLE-at='speccer -cmd meta -sec FEATURE -name TITLE -at';

alias spec-NONGOAL-RESPONSIBLE-at='speccer -cmd meta -sec NONGOAL -name RESPONSIBLE -at';
alias spec-NONGOAL-STATE-at='speccer -cmd meta -sec NONGOAL -name STATE -at';
alias spec-NONGOAL-DEADLINE-at='speccer -cmd meta -sec NONGOAL -name DEADLINE -at';
alias spec-NONGOAL-LASTUPDATE-at='speccer -cmd meta -sec NONGOAL -name LASTUPDATE -at';
alias spec-NONGOAL-ESTIMATEDHOURS-at='speccer -cmd meta -sec NONGOAL -name ESTIMATEDHOURS -at';
alias spec-NONGOAL-TITLE-at='speccer -cmd meta -sec NONGOAL -name TITLE -at';

alias spec-UNDECIDED-RESPONSIBLE-at='speccer -cmd meta -sec UNDECIDED -name RESPONSIBLE -at';
alias spec-UNDECIDED-STATE-at='speccer -cmd meta -sec UNDECIDED -name STATE -at';
alias spec-UNDECIDED-DEADLINE-at='speccer -cmd meta -sec UNDECIDED -name DEADLINE -at';
alias spec-UNDECIDED-LASTUPDATE-at='speccer -cmd meta -sec UNDECIDED -name LASTUPDATE -at';
alias spec-UNDECIDED-ESTIMATEDHOURS-at='speccer -cmd meta -sec UNDECIDED -name ESTIMATEDHOURS -at';
alias spec-UNDECIDED-TITLE-at='speccer -cmd meta -sec UNDECIDED -name TITLE -at';

# Set Meta Data from Section
alias spec-set-RESPONSIBLE='speccer -cmd meta -sec OVERVIEW -name RESPONSIBLE -set';
alias spec-set-PLANNING='speccer -cmd meta -sec OVERVIEW -name STATE -set PLANNING';
alias spec-set-APPROVED='speccer -cmd meta -sec OVERVIEW -name STATE -set APPROVED';
alias spec-set-PARTLY_IMPLEMENTED='speccer -cmd meta -sec OVERVIEW -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-FULLY_IMPLEMENTED='speccer -cmd meta -sec OVERVIEW -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-OBSOLET='speccer -cmd meta -sec OVERVIEW -name STATE -set OBSOLET';
alias spec-set-DEADLINE='speccer -cmd meta -sec OVERVIEW -name DEADLINE -set';
alias spec-set-ESTIMATEDHOURS='speccer -cmd meta -sec OVERVIEW -name ESTIMATEDHOURS -set';
alias spec-set-TITLE='speccer -cmd meta -sec OVERVIEW -name TITLE -set';

alias spec-set-SCENARIO-RESPONSIBLE='speccer -cmd meta -sec SCENARIO -name RESPONSIBLE -set';
alias spec-set-SCENARIO-PLANNING='speccer -cmd meta -sec SCENARIO -name STATE -set PLANNING';
alias spec-set-SCENARIO-APPROVED='speccer -cmd meta -sec SCENARIO -name STATE -set APPROVED';
alias spec-set-SCENARIO-PARTLY_IMPLEMENTED='speccer -cmd meta -sec SCENARIO -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-SCENARIO-FULLY_IMPLEMENTED='speccer -cmd meta -sec SCENARIO -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-SCENARIO-OBSOLET='speccer -cmd meta -sec SCENARIO -name STATE -set OBSOLET';
alias spec-set-SCENARIO-DEADLINE='speccer -cmd meta -sec SCENARIO -name DEADLINE -set';
alias spec-set-SCENARIO-ESTIMATEDHOURS='speccer -cmd meta -sec SCENARIO -name ESTIMATEDHOURS -set';
alias spec-set-SCENARIO-TITLE='speccer -cmd meta -sec SCENARIO -name TITLE -set';

alias spec-set-CONTRADICTION-RESPONSIBLE='speccer -cmd meta -sec CONTRADICTION -name RESPONSIBLE -set';
alias spec-set-CONTRADICTION-PLANNING='speccer -cmd meta -sec CONTRADICTION -name STATE -set PLANNING';
alias spec-set-CONTRADICTION-APPROVED='speccer -cmd meta -sec CONTRADICTION -name STATE -set APPROVED';
alias spec-set-CONTRADICTION-PARTLY_IMPLEMENTED='speccer -cmd meta -sec CONTRADICTION -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-CONTRADICTION-FULLY_IMPLEMENTED='speccer -cmd meta -sec CONTRADICTION -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-CONTRADICTION-OBSOLET='speccer -cmd meta -sec CONTRADICTION -name STATE -set OBSOLET';
alias spec-set-CONTRADICTION-DEADLINE='speccer -cmd meta -sec CONTRADICTION -name DEADLINE -set';
alias spec-set-CONTRADICTION-ESTIMATEDHOURS='speccer -cmd meta -sec CONTRADICTION -name ESTIMATEDHOURS -set';
alias spec-set-CONTRADICTION-TITLE='speccer -cmd meta -sec CONTRADICTION -name TITLE -set';

alias spec-set-DEFINITION-RESPONSIBLE='speccer -cmd meta -sec DEFINITION -name RESPONSIBLE -set';
alias spec-set-DEFINITION-PLANNING='speccer -cmd meta -sec DEFINITION -name STATE -set PLANNING';
alias spec-set-DEFINITION-APPROVED='speccer -cmd meta -sec DEFINITION -name STATE -set APPROVED';
alias spec-set-DEFINITION-PARTLY_IMPLEMENTED='speccer -cmd meta -sec DEFINITION -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-DEFINITION-FULLY_IMPLEMENTED='speccer -cmd meta -sec DEFINITION -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-DEFINITION-OBSOLET='speccer -cmd meta -sec DEFINITION -name STATE -set OBSOLET';
alias spec-set-DEFINITION-DEADLINE='speccer -cmd meta -sec DEFINITION -name DEADLINE -set';
alias spec-set-DEFINITION-ESTIMATEDHOURS='speccer -cmd meta -sec DEFINITION -name ESTIMATEDHOURS -set';
alias spec-set-DEFINITION-TITLE='speccer -cmd meta -sec DEFINITION -name TITLE -set';

alias spec-set-FEATURE-RESPONSIBLE='speccer -cmd meta -sec FEATURE -name RESPONSIBLE -set';
alias spec-set-FEATURE-PLANNING='speccer -cmd meta -sec FEATURE -name STATE -set PLANNING';
alias spec-set-FEATURE-APPROVED='speccer -cmd meta -sec FEATURE -name STATE -set APPROVED';
alias spec-set-FEATURE-PARTLY_IMPLEMENTED='speccer -cmd meta -sec FEATURE -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-FEATURE-FULLY_IMPLEMENTED='speccer -cmd meta -sec FEATURE -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-FEATURE-OBSOLET='speccer -cmd meta -sec FEATURE -name STATE -set OBSOLET';
alias spec-set-FEATURE-DEADLINE='speccer -cmd meta -sec FEATURE -name DEADLINE -set';
alias spec-set-FEATURE-ESTIMATEDHOURS='speccer -cmd meta -sec FEATURE -name ESTIMATEDHOURS -set';
alias spec-set-FEATURE-TITLE='speccer -cmd meta -sec FEATURE -name TITLE -set';

alias spec-set-NONGOAL-RESPONSIBLE='speccer -cmd meta -sec NONGOAL -name RESPONSIBLE -set';
alias spec-set-NONGOAL-PLANNING='speccer -cmd meta -sec NONGOAL -name STATE -set PLANNING';
alias spec-set-NONGOAL-APPROVED='speccer -cmd meta -sec NONGOAL -name STATE -set APPROVED';
alias spec-set-NONGOAL-PARTLY_IMPLEMENTED='speccer -cmd meta -sec NONGOAL -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-NONGOAL-FULLY_IMPLEMENTED='speccer -cmd meta -sec NONGOAL -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-NONGOAL-OBSOLET='speccer -cmd meta -sec NONGOAL -name STATE -set OBSOLET';
alias spec-set-NONGOAL-DEADLINE='speccer -cmd meta -sec NONGOAL -name DEADLINE -set';
alias spec-set-NONGOAL-ESTIMATEDHOURS='speccer -cmd meta -sec NONGOAL -name ESTIMATEDHOURS -set';
alias spec-set-NONGOAL-TITLE='speccer -cmd meta -sec NONGOAL -name TITLE -set';

alias spec-set-UNDECIDED-RESPONSIBLE='speccer -cmd meta -sec UNDECIDED -name RESPONSIBLE -set';
alias spec-set-UNDECIDED-PLANNING='speccer -cmd meta -sec UNDECIDED -name STATE -set PLANNING';
alias spec-set-UNDECIDED-APPROVED='speccer -cmd meta -sec UNDECIDED -name STATE -set APPROVED';
alias spec-set-UNDECIDED-PARTLY_IMPLEMENTED='speccer -cmd meta -sec UNDECIDED -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-UNDECIDED-FULLY_IMPLEMENTED='speccer -cmd meta -sec UNDECIDED -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-UNDECIDED-OBSOLET='speccer -cmd meta -sec UNDECIDED -name STATE -set OBSOLET';
alias spec-set-UNDECIDED-DEADLINE='speccer -cmd meta -sec UNDECIDED -name DEADLINE -set';
alias spec-set-UNDECIDED-ESTIMATEDHOURS='speccer -cmd meta -sec UNDECIDED -name ESTIMATEDHOURS -set';
alias spec-set-UNDECIDED-TITLE='speccer -cmd meta -sec UNDECIDED -name TITLE -set';

