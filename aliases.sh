# Common
alias spec_create='speccer -cmd create';
alias spec_validate='speccer -cmd validate';
alias spec_copy='speccer -cmd copy';

alias spec_markdown='speccer -cmd markdown';
alias spec_markdown-doc='speccer -cmd markdown -usage "documentation"';
alias spec_markdown-approve='speccer -cmd markdown -usage "approval"';
alias spec_markdown-discuss='speccer -cmd markdown -usage "discussion"';
alias spec_markdown-implement='speccer -cmd markdown -usage "implementation"';
alias spec_markdown-clean='speccer -cmd markdown -filter '"'"'{"META": true, "COMMENTS": true}'"'"'';

alias spec_html='speccer -cmd html';
alias spec_html-doc='speccer -cmd html -usage "documentation"';
alias spec_html-approve='speccer -cmd html -usage "approval"';
alias spec_html-discuss='speccer -cmd html -usage "discussion"';
alias spec_html-implement='speccer -cmd html -usage "implementation"';
alias spec_html-clean='speccer -cmd html -filter '"'"'{"META": true, "COMMENTS": true}'"'"'';

# Spec property getters
alias spec-requestedby='speccer -cmd property -name REQUESTEDBY';  
alias spec-related='speccer -cmd property -name RELATED';
alias spec-translations='speccer -cmd property -name TRANSLATIONS'; 
alias spec-supersededby='speccer -cmd property -name SUPERSEDEDBY'; 
alias spec-resources='speccer -cmd property -name RESOURCES';
alias spec-parent='speccer -cmd property -name PARENT';
alias spec-company='speccer -cmd property -name COMPANY';
alias spec-project='speccer -cmd property -name PROJECT';
alias spec-persons='speccer -cmd property -name PERSONS';
alias spec-url='speccer -cmd property -name URL';
alias spec-language='speccer -cmd property -name LANGUAGE';
alias spec-dateformat='speccer -cmd property -name DATEFORMAT';
alias spec-approved='speccer -cmd property -name APPROVED';     

# Spec property unsetters
alias spec_uset-requestedby='speccer -cmd property -name REQUESTEDBY -uset';  
alias spec_uset-related='speccer -cmd property -name RELATED -uset';
alias spec_uset-translations='speccer -cmd property -name TRANSLATIONS -uset'; 
alias spec_uset-supersededby='speccer -cmd property -name SUPERSEDEDBY -uset'; 
alias spec_uset-resources='speccer -cmd property -name RESOURCES -uset';
alias spec_uset-parent='speccer -cmd property -name PARENT -uset';
alias spec_uset-persons='speccer -cmd property -name PERSONS -uset';

# Spec property setters
alias spec_set-requestedby='speccer -cmd property -name REQUESTEDBY -set';  
alias spec_set-related='speccer -cmd property -name RELATED -set';
alias spec_set-translations='speccer -cmd property -name TRANSLATIONS -set'; 
alias spec_set-supersededby='speccer -cmd property -name SUPERSEDEDBY -set'; 
alias spec_set-resources='speccer -cmd property -name RESOURCES -set';
alias spec_set-parent='speccer -cmd property -name PARENT -set';
alias spec_set-company='speccer -cmd property -name COMPANY -set';
alias spec_set-project='speccer -cmd property -name PROJECT -set';
alias spec_set-url='speccer -cmd property -name URL -set';
alias spec_set-language='speccer -cmd property -name LANGUAGE -set';
alias spec_set-dateformat='speccer -cmd property -name DATEFORMAT -set';
alias spec_set-approved='speccer -cmd property -name APPROVED -set';
alias spec_set-persons='speccer -cmd property -name PERSONS -set';     

# Add Sections
alias spec_add-SCENARIO="speccer -cmd add -sec SCENARIO -resp $USER -set";
alias spec_add-CONTRADICTION="speccer -cmd add -sec CONTRADICTION -resp $USER -set";
alias spec_add-DEFINITION="speccer -cmd add -sec DEFINITION -resp $USER -set";
alias spec_add-FEATURE="speccer -cmd add -sec FEATURE -resp $USER -set";
alias spec_add-NONGOAL="speccer -cmd add -sec NONGOAL -resp $USER -set";
alias spec_add-UNDECIDED="speccer -cmd add -sec UNDECIDED -resp $USER -set";

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
alias spec_ls='speccer -cmd positions';
alias spec_ls-SCENARIO='speccer -cmd positions -sec SCENARIO';
alias spec_ls-CONTRADICTION='speccer -cmd positions -sec CONTRADICTION';
alias spec_ls-DEFINITION='speccer -cmd positions -sec DEFINITION';
alias spec_ls-FEATURE='speccer -cmd positions -sec FEATURE';
alias spec_ls-NONGOAL='speccer -cmd positions -sec NONGOAL';
alias spec_ls-UNDECIDED='speccer -cmd positions -sec UNDECIDED';

# Get uuids of Paragraphs for a section
alias spec_uuids='speccer -cmd uuids';
alias spec_uuids-SCENARIO='speccer -cmd uuids -sec SCENARIO';
alias spec_uuids-CONTRADICTION='speccer -cmd uuids -sec CONTRADICTION';
alias spec_uuids-DEFINITION='speccer -cmd uuids -sec DEFINITION';
alias spec_uuids-FEATURE='speccer -cmd uuids -sec FEATURE';
alias spec_uuids-NONGOAL='speccer -cmd uuids -sec NONGOAL';
alias spec_uuids-UNDECIDED='speccer -cmd uuids -sec UNDECIDED';

# Get Comments of Paragraphs for a section
alias spec_comment-OVERVIEW='speccer -cmd comment -sec OVERVIEW';
alias spec_comment-SCENARIO-at='speccer -cmd comment -sec SCENARIO -at';
alias spec_comment-CONTRADICTION-at='speccer -cmd comment -sec CONTRADICTION -at';
alias spec_comment-DEFINITION-at='speccer -cmd comment -sec DEFINITION -at';
alias spec_comment-FEATURE-at='speccer -cmd comment -sec FEATURE -at';
alias spec_comment-NONGOAL-at='speccer -cmd comment -sec NONGOAL -at';
alias spec_comment-UNDECIDED-at='speccer -cmd comment -sec UNDECIDED -at';

# Set Comment of Paragraph for a section
alias spec_comment_replace-OVERVIEW='speccer -cmd comment -sec OVERVIEW -set';
alias spec_comment_replace-SCENARIO='speccer -cmd comment -sec SCENARIO -set';
alias spec_comment_replace-CONTRADICTION='speccer -cmd comment -sec CONTRADICTION -set';
alias spec_comment_replace-DEFINITION='speccer -cmd comment -sec DEFINITION -set';
alias spec_comment_replace-FEATURE='speccer -cmd comment -sec FEATURE -set';
alias spec_comment_replace-NONGOAL='speccer -cmd comment -sec NONGOAL -set';
alias spec_comment_replace-UNDECIDED='speccer -cmd comment -sec UNDECIDED -set';

# Replace Sections
alias spec_replace-OVERVIEW='speccer -cmd text -sec OVERVIEW -set';
alias spec_replace-SCENARIO='speccer -cmd text -sec SCENARIO -set';
alias spec_replace-CONTRADICTION='speccer -cmd text -sec CONTRADICTION -set';
alias spec_replace-DEFINITION='speccer -cmd text -sec DEFINITION -set';
alias spec_replace-FEATURE='speccer -cmd text -sec FEATURE -set';
alias spec_replace-NONGOAL='speccer -cmd text -sec NONGOAL -set';
alias spec_replace-UNDECIDED='speccer -cmd text -sec UNDECIDED -set';

# Move Sections
alias spec_mv-SCENARIO-at='speccer -cmd move -sec SCENARIO -at';
alias spec_mv-CONTRADICTION-at='speccer -cmd move -sec CONTRADICTION -at';
alias spec_mv-DEFINITION-at='speccer -cmd move -sec DEFINITION -at';
alias spec_mv-FEATURE-at='speccer -cmd move -sec FEATURE -at';
alias spec_mv-NONGOAL-at='speccer -cmd move -sec NONGOAL -at';
alias spec_mv-UNDECIDED-at='speccer -cmd move -sec UNDECIDED -at';

# Remove Sections
alias spec_rm-SCENARIO-at='speccer -cmd rm -sec SCENARIO -at';
alias spec_rm-CONTRADICTION-at='speccer -cmd rm -sec CONTRADICTION -at';
alias spec_rm-DEFINITION-at='speccer -cmd rm -sec DEFINITION -at';
alias spec_rm-FEATURE-at='speccer -cmd rm -sec FEATURE -at';
alias spec_rm-NONGOAL-at='speccer -cmd rm -sec NONGOAL -at';
alias spec_rm-UNDECIDED-at='speccer -cmd rm -sec UNDECIDED -at';

# Get Meta Data from Section
alias spec-RESPONSIBLE='speccer -cmd meta -sec OVERVIEW -name RESPONSIBLE';
alias spec-STATE='speccer -cmd meta -sec OVERVIEW -name STATE';
alias spec-DEADLINE='speccer -cmd meta -sec OVERVIEW -name DEADLINE';
alias spec-LASTUPDATE='speccer -cmd meta -sec OVERVIEW -name LASTUPDATE';
alias spec-ESTIMATEDHOURS='speccer -cmd meta -sec OVERVIEW -name ESTIMATEDHOURS';
alias spec-TITLE='speccer -cmd meta -sec OVERVIEW -name TITLE';
alias spec-UUID='speccer -cmd meta -sec OVERVIEW -name UUID';

alias spec-SCENARIO-RESPONSIBLE-at='speccer -cmd meta -sec SCENARIO -name RESPONSIBLE -at';
alias spec-SCENARIO-STATE-at='speccer -cmd meta -sec SCENARIO -name STATE -at';
alias spec-SCENARIO-DEADLINE-at='speccer -cmd meta -sec SCENARIO -name DEADLINE -at';
alias spec-SCENARIO-LASTUPDATE-at='speccer -cmd meta -sec SCENARIO -name LASTUPDATE -at';
alias spec-SCENARIO-ESTIMATEDHOURS-at='speccer -cmd meta -sec SCENARIO -name ESTIMATEDHOURS -at';
alias spec-SCENARIO-TITLE-at='speccer -cmd meta -sec SCENARIO -name TITLE -at';
alias spec-SCENARIO-UUID-at='speccer -cmd meta -sec SCENARIO -name UUID -at';

alias spec-CONTRADICTION-RESPONSIBLE-at='speccer -cmd meta -sec CONTRADICTION -name RESPONSIBLE -at';
alias spec-CONTRADICTION-STATE-at='speccer -cmd meta -sec CONTRADICTION -name STATE -at';
alias spec-CONTRADICTION-DEADLINE-at='speccer -cmd meta -sec CONTRADICTION -name DEADLINE -at';
alias spec-CONTRADICTION-LASTUPDATE-at='speccer -cmd meta -sec CONTRADICTION -name LASTUPDATE -at';
alias spec-CONTRADICTION-ESTIMATEDHOURS-at='speccer -cmd meta -sec CONTRADICTION -name ESTIMATEDHOURS -at';
alias spec-CONTRADICTION-TITLE-at='speccer -cmd meta -sec CONTRADICTION -name TITLE -at';
alias spec-CONTRADICTION-UUID-at='speccer -cmd meta -sec CONTRADICTION -name UUID -at';

alias spec-DEFINITION-RESPONSIBLE-at='speccer -cmd meta -sec DEFINITION -name RESPONSIBLE -at';
alias spec-DEFINITION-STATE-at='speccer -cmd meta -sec DEFINITION -name STATE -at';
alias spec-DEFINITION-DEADLINE-at='speccer -cmd meta -sec DEFINITION -name DEADLINE -at';
alias spec-DEFINITION-LASTUPDATE-at='speccer -cmd meta -sec DEFINITION -name LASTUPDATE -at';
alias spec-DEFINITION-ESTIMATEDHOURS-at='speccer -cmd meta -sec DEFINITION -name ESTIMATEDHOURS -at';
alias spec-DEFINITION-TITLE-at='speccer -cmd meta -sec DEFINITION -name TITLE -at';
alias spec-DEFINITION-UUID-at='speccer -cmd meta -sec DEFINITION -name UUID -at';

alias spec-FEATURE-RESPONSIBLE-at='speccer -cmd meta -sec FEATURE -name RESPONSIBLE -at';
alias spec-FEATURE-STATE-at='speccer -cmd meta -sec FEATURE -name STATE -at';
alias spec-FEATURE-DEADLINE-at='speccer -cmd meta -sec FEATURE -name DEADLINE -at';
alias spec-FEATURE-LASTUPDATE-at='speccer -cmd meta -sec FEATURE -name LASTUPDATE -at';
alias spec-FEATURE-ESTIMATEDHOURS-at='speccer -cmd meta -sec FEATURE -name ESTIMATEDHOURS -at';
alias spec-FEATURE-TITLE-at='speccer -cmd meta -sec FEATURE -name TITLE -at';
alias spec-FEATURE-UUID-at='speccer -cmd meta -sec FEATURE -name UUID -at';

alias spec-NONGOAL-RESPONSIBLE-at='speccer -cmd meta -sec NONGOAL -name RESPONSIBLE -at';
alias spec-NONGOAL-STATE-at='speccer -cmd meta -sec NONGOAL -name STATE -at';
alias spec-NONGOAL-DEADLINE-at='speccer -cmd meta -sec NONGOAL -name DEADLINE -at';
alias spec-NONGOAL-LASTUPDATE-at='speccer -cmd meta -sec NONGOAL -name LASTUPDATE -at';
alias spec-NONGOAL-ESTIMATEDHOURS-at='speccer -cmd meta -sec NONGOAL -name ESTIMATEDHOURS -at';
alias spec-NONGOAL-TITLE-at='speccer -cmd meta -sec NONGOAL -name TITLE -at';
alias spec-NONGOAL-UUID-at='speccer -cmd meta -sec NONGOAL -name UUID -at';

alias spec-UNDECIDED-RESPONSIBLE-at='speccer -cmd meta -sec UNDECIDED -name RESPONSIBLE -at';
alias spec-UNDECIDED-STATE-at='speccer -cmd meta -sec UNDECIDED -name STATE -at';
alias spec-UNDECIDED-DEADLINE-at='speccer -cmd meta -sec UNDECIDED -name DEADLINE -at';
alias spec-UNDECIDED-LASTUPDATE-at='speccer -cmd meta -sec UNDECIDED -name LASTUPDATE -at';
alias spec-UNDECIDED-ESTIMATEDHOURS-at='speccer -cmd meta -sec UNDECIDED -name ESTIMATEDHOURS -at';
alias spec-UNDECIDED-TITLE-at='speccer -cmd meta -sec UNDECIDED -name TITLE -at';
alias spec-NONGOAL-UUID-at='speccer -cmd meta -sec NONGOAL -name UUID -at';

# Set Meta Data from Section
alias spec_set-RESPONSIBLE='speccer -cmd meta -sec OVERVIEW -name RESPONSIBLE -set';
alias spec_set-PLANNING='speccer -cmd meta -sec OVERVIEW -name STATE -set PLANNING';
alias spec_set-APPROVED='speccer -cmd meta -sec OVERVIEW -name STATE -set APPROVED';
alias spec_set-PARTLY_IMPLEMENTED='speccer -cmd meta -sec OVERVIEW -name STATE -set PARTLY_IMPLEMENTED';
alias spec_set-FULLY_IMPLEMENTED='speccer -cmd meta -sec OVERVIEW -name STATE -set FULLY_IMPLEMENTED';
alias spec_set-OBSOLET='speccer -cmd meta -sec OVERVIEW -name STATE -set OBSOLET';
alias spec_set-DEADLINE='speccer -cmd meta -sec OVERVIEW -name DEADLINE -set';
alias spec_set-ESTIMATEDHOURS='speccer -cmd meta -sec OVERVIEW -name ESTIMATEDHOURS -set';
alias spec_set-TITLE='speccer -cmd meta -sec OVERVIEW -name TITLE -set';

alias spec_set-SCENARIO-RESPONSIBLE='speccer -cmd meta -sec SCENARIO -name RESPONSIBLE -set';
alias spec_set-SCENARIO-PLANNING='speccer -cmd meta -sec SCENARIO -name STATE -set PLANNING';
alias spec_set-SCENARIO-APPROVED='speccer -cmd meta -sec SCENARIO -name STATE -set APPROVED';
alias spec_set-SCENARIO-PARTLY_IMPLEMENTED='speccer -cmd meta -sec SCENARIO -name STATE -set PARTLY_IMPLEMENTED';
alias spec_set-SCENARIO-FULLY_IMPLEMENTED='speccer -cmd meta -sec SCENARIO -name STATE -set FULLY_IMPLEMENTED';
alias spec_set-SCENARIO-OBSOLET='speccer -cmd meta -sec SCENARIO -name STATE -set OBSOLET';
alias spec_set-SCENARIO-DEADLINE='speccer -cmd meta -sec SCENARIO -name DEADLINE -set';
alias spec_set-SCENARIO-ESTIMATEDHOURS='speccer -cmd meta -sec SCENARIO -name ESTIMATEDHOURS -set';
alias spec_set-SCENARIO-TITLE='speccer -cmd meta -sec SCENARIO -name TITLE -set';

alias spec_set-CONTRADICTION-RESPONSIBLE='speccer -cmd meta -sec CONTRADICTION -name RESPONSIBLE -set';
alias spec_set-CONTRADICTION-PLANNING='speccer -cmd meta -sec CONTRADICTION -name STATE -set PLANNING';
alias spec_set-CONTRADICTION-APPROVED='speccer -cmd meta -sec CONTRADICTION -name STATE -set APPROVED';
alias spec_set-CONTRADICTION-PARTLY_IMPLEMENTED='speccer -cmd meta -sec CONTRADICTION -name STATE -set PARTLY_IMPLEMENTED';
alias spec_set-CONTRADICTION-FULLY_IMPLEMENTED='speccer -cmd meta -sec CONTRADICTION -name STATE -set FULLY_IMPLEMENTED';
alias spec_set-CONTRADICTION-OBSOLET='speccer -cmd meta -sec CONTRADICTION -name STATE -set OBSOLET';
alias spec_set-CONTRADICTION-DEADLINE='speccer -cmd meta -sec CONTRADICTION -name DEADLINE -set';
alias spec_set-CONTRADICTION-ESTIMATEDHOURS='speccer -cmd meta -sec CONTRADICTION -name ESTIMATEDHOURS -set';
alias spec_set-CONTRADICTION-TITLE='speccer -cmd meta -sec CONTRADICTION -name TITLE -set';

alias spec_set-DEFINITION-RESPONSIBLE='speccer -cmd meta -sec DEFINITION -name RESPONSIBLE -set';
alias spec_set-DEFINITION-PLANNING='speccer -cmd meta -sec DEFINITION -name STATE -set PLANNING';
alias spec_set-DEFINITION-APPROVED='speccer -cmd meta -sec DEFINITION -name STATE -set APPROVED';
alias spec_set-DEFINITION-PARTLY_IMPLEMENTED='speccer -cmd meta -sec DEFINITION -name STATE -set PARTLY_IMPLEMENTED';
alias spec_set-DEFINITION-FULLY_IMPLEMENTED='speccer -cmd meta -sec DEFINITION -name STATE -set FULLY_IMPLEMENTED';
alias spec_set-DEFINITION-OBSOLET='speccer -cmd meta -sec DEFINITION -name STATE -set OBSOLET';
alias spec_set-DEFINITION-DEADLINE='speccer -cmd meta -sec DEFINITION -name DEADLINE -set';
alias spec_set-DEFINITION-ESTIMATEDHOURS='speccer -cmd meta -sec DEFINITION -name ESTIMATEDHOURS -set';
alias spec_set-DEFINITION-TITLE='speccer -cmd meta -sec DEFINITION -name TITLE -set';

alias spec_set-FEATURE-RESPONSIBLE='speccer -cmd meta -sec FEATURE -name RESPONSIBLE -set';
alias spec_set-FEATURE-PLANNING='speccer -cmd meta -sec FEATURE -name STATE -set PLANNING';
alias spec_set-FEATURE-APPROVED='speccer -cmd meta -sec FEATURE -name STATE -set APPROVED';
alias spec_set-FEATURE-PARTLY_IMPLEMENTED='speccer -cmd meta -sec FEATURE -name STATE -set PARTLY_IMPLEMENTED';
alias spec_set-FEATURE-FULLY_IMPLEMENTED='speccer -cmd meta -sec FEATURE -name STATE -set FULLY_IMPLEMENTED';
alias spec_set-FEATURE-OBSOLET='speccer -cmd meta -sec FEATURE -name STATE -set OBSOLET';
alias spec_set-FEATURE-DEADLINE='speccer -cmd meta -sec FEATURE -name DEADLINE -set';
alias spec_set-FEATURE-ESTIMATEDHOURS='speccer -cmd meta -sec FEATURE -name ESTIMATEDHOURS -set';
alias spec_set-FEATURE-TITLE='speccer -cmd meta -sec FEATURE -name TITLE -set';

alias spec_set-NONGOAL-RESPONSIBLE='speccer -cmd meta -sec NONGOAL -name RESPONSIBLE -set';
alias spec_set-NONGOAL-PLANNING='speccer -cmd meta -sec NONGOAL -name STATE -set PLANNING';
alias spec_set-NONGOAL-APPROVED='speccer -cmd meta -sec NONGOAL -name STATE -set APPROVED';
alias spec_set-NONGOAL-PARTLY_IMPLEMENTED='speccer -cmd meta -sec NONGOAL -name STATE -set PARTLY_IMPLEMENTED';
alias spec_set-NONGOAL-FULLY_IMPLEMENTED='speccer -cmd meta -sec NONGOAL -name STATE -set FULLY_IMPLEMENTED';
alias spec_set-NONGOAL-OBSOLET='speccer -cmd meta -sec NONGOAL -name STATE -set OBSOLET';
alias spec_set-NONGOAL-DEADLINE='speccer -cmd meta -sec NONGOAL -name DEADLINE -set';
alias spec_set-NONGOAL-ESTIMATEDHOURS='speccer -cmd meta -sec NONGOAL -name ESTIMATEDHOURS -set';
alias spec_set-NONGOAL-TITLE='speccer -cmd meta -sec NONGOAL -name TITLE -set';

alias spec_set-UNDECIDED-RESPONSIBLE='speccer -cmd meta -sec UNDECIDED -name RESPONSIBLE -set';
alias spec_set-UNDECIDED-PLANNING='speccer -cmd meta -sec UNDECIDED -name STATE -set PLANNING';
alias spec_set-UNDECIDED-APPROVED='speccer -cmd meta -sec UNDECIDED -name STATE -set APPROVED';
alias spec_set-UNDECIDED-PARTLY_IMPLEMENTED='speccer -cmd meta -sec UNDECIDED -name STATE -set PARTLY_IMPLEMENTED';
alias spec_set-UNDECIDED-FULLY_IMPLEMENTED='speccer -cmd meta -sec UNDECIDED -name STATE -set FULLY_IMPLEMENTED';
alias spec_set-UNDECIDED-OBSOLET='speccer -cmd meta -sec UNDECIDED -name STATE -set OBSOLET';
alias spec_set-UNDECIDED-DEADLINE='speccer -cmd meta -sec UNDECIDED -name DEADLINE -set';
alias spec_set-UNDECIDED-ESTIMATEDHOURS='speccer -cmd meta -sec UNDECIDED -name ESTIMATEDHOURS -set';
alias spec_set-UNDECIDED-TITLE='speccer -cmd meta -sec UNDECIDED -name TITLE -set';

