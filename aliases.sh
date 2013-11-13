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
alias spec-TITLE='speccer -cmd property -name TITLE';
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
alias spec-TITLE-set='speccer -cmd property -name TITLE -set';
alias spec-URL-set='speccer -cmd property -name URL -set';
alias spec-LANGUAGE-set='speccer -cmd property -name LANGUAGE -set';
alias spec-DATEFORMAT-set='speccer -cmd property -name DATEFORMAT -set';
alias spec-APPROVED-set='speccer -cmd property -name APPROVED -set';     
alias spec-PERSONS-set='speccer -cmd property -name PERSONS -set';     

# Add Sections
alias spec-add-SCENARIO="speccer -cmd add -sec SCENARIOS -resp $USER -set";
alias spec-add-CONTRADICTION="speccer -cmd add -sec CONTRADICTIONS -resp $USER -set";
alias spec-add-DEFINITION="speccer -cmd add -sec DEFINITIONS -resp $USER -set";
alias spec-add-FEATURE="speccer -cmd add -sec FEATURES -resp $USER -set";
alias spec-add-NONGOAL="speccer -cmd add -sec NONGOALS -resp $USER -set";
alias spec-add-UNDECIDED="speccer -cmd add -sec UNDECIDED -resp $USER -set";

# Show Sections
alias spec-OVERVIEW='speccer -cmd text -sec OVERVIEW';
alias spec-SCENARIO-at='speccer -cmd text -sec SCENARIOS -at';
alias spec-CONTRADICTION-at='speccer -cmd text -sec CONTRADICTIONS -at';
alias spec-DEFINITION-at='speccer -cmd text -sec DEFINITIONS -at';
alias spec-FEATURE-at='speccer -cmd text -sec FEATURES -at';
alias spec-NONGOAL-at='speccer -cmd text -sec NONGOALS -at';
alias spec-UNDECIDED-at='speccer -cmd text -sec UNDECIDED -at';

alias spec-OVERVIEW-full='speccer -cmd text -sec OVERVIEW -with-comments -with-meta';
alias spec-SCENARIO-full-at='speccer -cmd text -sec SCENARIOS -with-comments -with-meta -at';
alias spec-CONTRADICTION-full-at='speccer -cmd text -sec CONTRADICTIONS -with-comments -with-meta -at';
alias spec-DEFINITION-full-at='speccer -cmd text -sec DEFINITIONS -with-comments -with-meta -at';
alias spec-FEATURE-full-at='speccer -cmd text -sec FEATURES -with-comments -with-meta -at';
alias spec-NONGOAL-full-at='speccer -cmd text -sec NONGOALS -with-comments -with-meta -at';
alias spec-UNDECIDED-full-at='speccer -cmd text -sec UNDECIDED -with-comments -with-meta -at';

# Get list of Paragraphs for a section
alias spec-ls='speccer -cmd positions';
alias spec-ls-SCENARIO='speccer -cmd positions -sec SCENARIOS';
alias spec-ls-CONTRADICTION='speccer -cmd positions -sec CONTRADICTIONS';
alias spec-ls-DEFINITION='speccer -cmd positions -sec DEFINITIONS';
alias spec-ls-FEATURE='speccer -cmd positions -sec FEATURES';
alias spec-ls-NONGOAL='speccer -cmd positions -sec NONGOALS';
alias spec-ls-UNDECIDED='speccer -cmd positions -sec UNDECIDED';

# Get Comments of Paragraphs for a section
alias spec-comment-OVERVIEW='speccer -cmd comment -sec OVERVIEW';
alias spec-comment-SCENARIO-at='speccer -cmd comment -sec SCENARIOS -at';
alias spec-comment-CONTRADICTION-at='speccer -cmd comment -sec CONTRADICTIONS -at';
alias spec-comment-DEFINITION-at='speccer -cmd comment -sec DEFINITIONS -at';
alias spec-comment-FEATURE-at='speccer -cmd comment -sec FEATURES -at';
alias spec-comment-NONGOAL-at='speccer -cmd comment -sec NONGOALS -at';
alias spec-comment-UNDECIDED-at='speccer -cmd comment -sec UNDECIDED -at';

# Set Comment of Paragraph for a section
alias spec-comment-OVERVIEW-set='speccer -cmd comment -sec OVERVIEW -set';
alias spec-comment-SCENARIO-set='speccer -cmd comment -sec SCENARIOS -set';
alias spec-comment-CONTRADICTION-set='speccer -cmd comment -sec CONTRADICTIONS -set';
alias spec-comment-DEFINITION-set='speccer -cmd comment -sec DEFINITIONS -set';
alias spec-comment-FEATURE-set='speccer -cmd comment -sec FEATURES -set';
alias spec-comment-NONGOAL-set='speccer -cmd comment -sec NONGOALS -set';
alias spec-comment-UNDECIDED-set='speccer -cmd comment -sec UNDECIDED -set';

# Replace Sections
alias spec-set-OVERVIEW='speccer -cmd text -sec OVERVIEW -set';
alias spec-set-SCENARIO='speccer -cmd text -sec SCENARIOS -set';
alias spec-set-CONTRADICTION='speccer -cmd text -sec CONTRADICTIONS -set';
alias spec-set-DEFINITION='speccer -cmd text -sec DEFINITIONS -set';
alias spec-set-FEATURE='speccer -cmd text -sec FEATURES -set';
alias spec-set-NONGOAL='speccer -cmd text -sec NONGOALS -set';
alias spec-set-UNDECIDED='speccer -cmd text -sec UNDECIDED -set';

# Move Sections
alias spec-mv-SCENARIO-at='speccer -cmd move -sec SCENARIOS -at';
alias spec-mv-CONTRADICTION-at='speccer -cmd move -sec CONTRADICTIONS -at';
alias spec-mv-DEFINITION-at='speccer -cmd move -sec DEFINITIONS -at';
alias spec-mv-FEATURE-at='speccer -cmd move -sec FEATURES -at';
alias spec-mv-NONGOAL-at='speccer -cmd move -sec NONGOALS -at';
alias spec-mv-UNDECIDED-at='speccer -cmd move -sec UNDECIDED -at';

# Remove Sections
alias spec-rm-SCENARIO-at='speccer -cmd rm -sec SCENARIOS -at';
alias spec-rm-CONTRADICTION-at='speccer -cmd rm -sec CONTRADICTIONS -at';
alias spec-rm-DEFINITION-at='speccer -cmd rm -sec DEFINITIONS -at';
alias spec-rm-FEATURE-at='speccer -cmd rm -sec FEATURES -at';
alias spec-rm-NONGOAL-at='speccer -cmd rm -sec NONGOALS -at';
alias spec-rm-UNDECIDED-at='speccer -cmd rm -sec UNDECIDED -at';

# Get Meta Data from Section
alias spec-RESPONSIBLE='speccer -cmd meta -sec OVERVIEW -name RESPONSIBLE';
alias spec-STATE='speccer -cmd meta -sec OVERVIEW -name STATE';
alias spec-DEADLINE='speccer -cmd meta -sec OVERVIEW -name DEADLINE';
alias spec-LASTUPDATE='speccer -cmd meta -sec OVERVIEW -name LASTUPDATE';
alias spec-ESTIMATEDHOURS='speccer -cmd meta -sec OVERVIEW -name ESTIMATEDHOURS';

alias spec-SCENARIO-RESPONSIBLE-at='speccer -cmd meta -sec SCENARIOS -name RESPONSIBLE -at';
alias spec-SCENARIO-STATE-at='speccer -cmd meta -sec SCENARIOS -name STATE -at';
alias spec-SCENARIO-DEADLINE-at='speccer -cmd meta -sec SCENARIOS -name DEADLINE -at';
alias spec-SCENARIO-LASTUPDATE-at='speccer -cmd meta -sec SCENARIOS -name LASTUPDATE -at';
alias spec-SCENARIO-ESTIMATEDHOURS-at='speccer -cmd meta -sec SCENARIOS -name ESTIMATEDHOURS -at';

alias spec-CONTRADICTION-RESPONSIBLE-at='speccer -cmd meta -sec CONTRADICTIONS -name RESPONSIBLE -at';
alias spec-CONTRADICTION-STATE-at='speccer -cmd meta -sec CONTRADICTIONS -name STATE -at';
alias spec-CONTRADICTION-DEADLINE-at='speccer -cmd meta -sec CONTRADICTIONS -name DEADLINE -at';
alias spec-CONTRADICTION-LASTUPDATE-at='speccer -cmd meta -sec CONTRADICTIONS -name LASTUPDATE -at';
alias spec-CONTRADICTION-ESTIMATEDHOURS-at='speccer -cmd meta -sec CONTRADICTIONS -name ESTIMATEDHOURS -at';

alias spec-DEFINITION-RESPONSIBLE-at='speccer -cmd meta -sec DEFINITIONS -name RESPONSIBLE -at';
alias spec-DEFINITION-STATE-at='speccer -cmd meta -sec DEFINITIONS -name STATE -at';
alias spec-DEFINITION-DEADLINE-at='speccer -cmd meta -sec DEFINITIONS -name DEADLINE -at';
alias spec-DEFINITION-LASTUPDATE-at='speccer -cmd meta -sec DEFINITIONS -name LASTUPDATE -at';
alias spec-DEFINITION-ESTIMATEDHOURS-at='speccer -cmd meta -sec DEFINITIONS -name ESTIMATEDHOURS -at';

alias spec-FEATURE-RESPONSIBLE-at='speccer -cmd meta -sec FEATURES -name RESPONSIBLE -at';
alias spec-FEATURE-STATE-at='speccer -cmd meta -sec FEATURES -name STATE -at';
alias spec-FEATURE-DEADLINE-at='speccer -cmd meta -sec FEATURES -name DEADLINE -at';
alias spec-FEATURE-LASTUPDATE-at='speccer -cmd meta -sec FEATURES -name LASTUPDATE -at';
alias spec-FEATURE-ESTIMATEDHOURS-at='speccer -cmd meta -sec FEATURES -name ESTIMATEDHOURS -at';

alias spec-NONGOAL-RESPONSIBLE-at='speccer -cmd meta -sec NONGOALS -name RESPONSIBLE -at';
alias spec-NONGOAL-STATE-at='speccer -cmd meta -sec NONGOALS -name STATE -at';
alias spec-NONGOAL-DEADLINE-at='speccer -cmd meta -sec NONGOALS -name DEADLINE -at';
alias spec-NONGOAL-LASTUPDATE-at='speccer -cmd meta -sec NONGOALS -name LASTUPDATE -at';
alias spec-NONGOAL-ESTIMATEDHOURS-at='speccer -cmd meta -sec NONGOALS -name ESTIMATEDHOURS -at';

alias spec-UNDECIDED-RESPONSIBLE-at='speccer -cmd meta -sec UNDECIDED -name RESPONSIBLE -at';
alias spec-UNDECIDED-STATE-at='speccer -cmd meta -sec UNDECIDED -name STATE -at';
alias spec-UNDECIDED-DEADLINE-at='speccer -cmd meta -sec UNDECIDED -name DEADLINE -at';
alias spec-UNDECIDED-LASTUPDATE-at='speccer -cmd meta -sec UNDECIDED -name LASTUPDATE -at';
alias spec-UNDECIDED-ESTIMATEDHOURS-at='speccer -cmd meta -sec UNDECIDED -name ESTIMATEDHOURS -at';

# Set Meta Data from Section
alias spec-set-RESPONSIBLE='speccer -cmd meta -sec OVERVIEW -name RESPONSIBLE -set';
alias spec-set-PLANNING='speccer -cmd meta -sec OVERVIEW -name STATE -set PLANNING';
alias spec-set-APPROVED='speccer -cmd meta -sec OVERVIEW -name STATE -set APPROVED';
alias spec-set-PARTLY_IMPLEMENTED='speccer -cmd meta -sec OVERVIEW -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-FULLY_IMPLEMENTED='speccer -cmd meta -sec OVERVIEW -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-OBSOLET='speccer -cmd meta -sec OVERVIEW -name STATE -set OBSOLET';
alias spec-set-DEADLINE='speccer -cmd meta -sec OVERVIEW -name DEADLINE -set';
alias spec-set-ESTIMATEDHOURS='speccer -cmd meta -sec OVERVIEW -name ESTIMATEDHOURS -set';

alias spec-set-SCENARIO-RESPONSIBLE='speccer -cmd meta -sec SCENARIOS -name RESPONSIBLE -set';
alias spec-set-SCENARIO-PLANNING='speccer -cmd meta -sec SCENARIOS -name STATE -set PLANNING';
alias spec-set-SCENARIO-APPROVED='speccer -cmd meta -sec SCENARIOS -name STATE -set APPROVED';
alias spec-set-SCENARIO-PARTLY_IMPLEMENTED='speccer -cmd meta -sec SCENARIOS -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-SCENARIO-FULLY_IMPLEMENTED='speccer -cmd meta -sec SCENARIOS -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-SCENARIO-OBSOLET='speccer -cmd meta -sec SCENARIOS -name STATE -set OBSOLET';
alias spec-set-SCENARIO-DEADLINE='speccer -cmd meta -sec SCENARIOS -name DEADLINE -set';
alias spec-set-SCENARIO-ESTIMATEDHOURS='speccer -cmd meta -sec SCENARIOS -name ESTIMATEDHOURS -set';

alias spec-set-CONTRADICTION-RESPONSIBLE='speccer -cmd meta -sec CONTRADICTIONS -name RESPONSIBLE -set';
alias spec-set-CONTRADICTION-PLANNING='speccer -cmd meta -sec CONTRADICTIONS -name STATE -set PLANNING';
alias spec-set-CONTRADICTION-APPROVED='speccer -cmd meta -sec CONTRADICTIONS -name STATE -set APPROVED';
alias spec-set-CONTRADICTION-PARTLY_IMPLEMENTED='speccer -cmd meta -sec CONTRADICTIONS -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-CONTRADICTION-FULLY_IMPLEMENTED='speccer -cmd meta -sec CONTRADICTIONS -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-CONTRADICTION-OBSOLET='speccer -cmd meta -sec CONTRADICTIONS -name STATE -set OBSOLET';
alias spec-set-CONTRADICTION-DEADLINE='speccer -cmd meta -sec CONTRADICTIONS -name DEADLINE -set';
alias spec-set-CONTRADICTION-ESTIMATEDHOURS='speccer -cmd meta -sec CONTRADICTIONS -name ESTIMATEDHOURS -set';

alias spec-set-DEFINITION-RESPONSIBLE='speccer -cmd meta -sec DEFINITIONS -name RESPONSIBLE -set';
alias spec-set-DEFINITION-PLANNING='speccer -cmd meta -sec DEFINITIONS -name STATE -set PLANNING';
alias spec-set-DEFINITION-APPROVED='speccer -cmd meta -sec DEFINITIONS -name STATE -set APPROVED';
alias spec-set-DEFINITION-PARTLY_IMPLEMENTED='speccer -cmd meta -sec DEFINITIONS -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-DEFINITION-FULLY_IMPLEMENTED='speccer -cmd meta -sec DEFINITIONS -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-DEFINITION-OBSOLET='speccer -cmd meta -sec DEFINITIONS -name STATE -set OBSOLET';
alias spec-set-DEFINITION-DEADLINE='speccer -cmd meta -sec DEFINITIONS -name DEADLINE -set';
alias spec-set-DEFINITION-ESTIMATEDHOURS='speccer -cmd meta -sec DEFINITIONS -name ESTIMATEDHOURS -set';

alias spec-set-FEATURE-RESPONSIBLE='speccer -cmd meta -sec FEATURES -name RESPONSIBLE -set';
alias spec-set-FEATURE-PLANNING='speccer -cmd meta -sec FEATURES -name STATE -set PLANNING';
alias spec-set-FEATURE-APPROVED='speccer -cmd meta -sec FEATURES -name STATE -set APPROVED';
alias spec-set-FEATURE-PARTLY_IMPLEMENTED='speccer -cmd meta -sec FEATURES -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-FEATURE-FULLY_IMPLEMENTED='speccer -cmd meta -sec FEATURES -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-FEATURE-OBSOLET='speccer -cmd meta -sec FEATURES -name STATE -set OBSOLET';
alias spec-set-FEATURE-DEADLINE='speccer -cmd meta -sec FEATURES -name DEADLINE -set';
alias spec-set-FEATURE-ESTIMATEDHOURS='speccer -cmd meta -sec FEATURES -name ESTIMATEDHOURS -set';

alias spec-set-NONGOAL-RESPONSIBLE='speccer -cmd meta -sec NONGOALS -name RESPONSIBLE -set';
alias spec-set-NONGOAL-PLANNING='speccer -cmd meta -sec NONGOALS -name STATE -set PLANNING';
alias spec-set-NONGOAL-APPROVED='speccer -cmd meta -sec NONGOALS -name STATE -set APPROVED';
alias spec-set-NONGOAL-PARTLY_IMPLEMENTED='speccer -cmd meta -sec NONGOALS -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-NONGOAL-FULLY_IMPLEMENTED='speccer -cmd meta -sec NONGOALS -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-NONGOAL-OBSOLET='speccer -cmd meta -sec NONGOALS -name STATE -set OBSOLET';
alias spec-set-NONGOAL-DEADLINE='speccer -cmd meta -sec NONGOALS -name DEADLINE -set';
alias spec-set-NONGOAL-ESTIMATEDHOURS='speccer -cmd meta -sec NONGOALS -name ESTIMATEDHOURS -set';

alias spec-set-UNDECIDED-RESPONSIBLE='speccer -cmd meta -sec UNDECIDED -name RESPONSIBLE -set';
alias spec-set-UNDECIDED-PLANNING='speccer -cmd meta -sec UNDECIDED -name STATE -set PLANNING';
alias spec-set-UNDECIDED-APPROVED='speccer -cmd meta -sec UNDECIDED -name STATE -set APPROVED';
alias spec-set-UNDECIDED-PARTLY_IMPLEMENTED='speccer -cmd meta -sec UNDECIDED -name STATE -set PARTLY_IMPLEMENTED';
alias spec-set-UNDECIDED-FULLY_IMPLEMENTED='speccer -cmd meta -sec UNDECIDED -name STATE -set FULLY_IMPLEMENTED';
alias spec-set-UNDECIDED-OBSOLET='speccer -cmd meta -sec UNDECIDED -name STATE -set OBSOLET';
alias spec-set-UNDECIDED-DEADLINE='speccer -cmd meta -sec UNDECIDED -name DEADLINE -set';
alias spec-set-UNDECIDED-ESTIMATEDHOURS='speccer -cmd meta -sec UNDECIDED -name ESTIMATEDHOURS -set';

