# How documentation is structured

Documentation is built to work on support.vorteil.io, which uses Kayako 
to structure and present information. Kayako breaks down documentation into 
three levels: categories, sections, and articles. We will mirror this breakdown
within this directory: each directory within this one is its own category,
each directory within a category directory is its own section. Each directory
within a section is its own article. Article directories should not contain 
other directories.

Categories and sections are designed to be no more than a single paragraph, 
and any description should convey as much information as possible within 
its first 120 characters, since Kayako will truncate it in some places.
To provide the description for a category or a section, create a README.md
file within its directory. README.md is a reserved file name for this purpose.
Despite the name, category and section README.md files should avoid using 
Markdown other than to provide a title due to the way Kayako sometimes presents
this information. 

Providing the content for an article is done similarly, by adding a README.md 
to an article directory. Article content follows very different guidelines.
Markdown is expected, and there is no need to limit the content to a single 
paragraph or blurb.

# Articles

## Title

Every article should begin with the page title as a markdown header (single 
hash). This title should match the directory name if possible. In the event 
where the directory name and the title do not match, the title within the 
contents will take precedence for display purposes but the title within the 
directory will be used for the purposes of linking articles. This is done so 
that typos and other reasons to adjust a title don't cause links to break.

## Keywords 

Kayako can improve an articles visibility in search results if we help it out by
providing likely keywords and phrases that users might search when they're 
looking for this article. This information can be provided by including a 
"keywords.md" file within an article directory with one keyword (or phrase) on 
each line.

## Links 

Articles are much more helpful if they provide links to other information. We
will support linking both externally and internally. If a Markdown link appears 
within an article and it begins with the protocol information ("http" or 
"https") then the link will be preserved. If the link does not contain protocol 
information it will be resolved by the uploading tool.

The uploading tool will split a link on its forward slashes and determine which 
category, section, or article to link to, resolving Kayako's URLs for these 
targets as necessary. 

## Attachments 

Some articles may like to add attachments. This could be particularly useful 
for guides. Any file within an article directory that is not one of the 
types of files otherwise named in this guide will be added to an article as an 
attachment.

## Misc 

The tools we use to sync these docs with the kayako portal currently have 
issues whenever the markdown files contain something that can be misinterpreted 
as HTML. Avoid using the following characters where possible for now: <>&.