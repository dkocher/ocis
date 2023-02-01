@api
Feature: create folder using MKCOL

  Background:
    Given user "Alice" has been created with default attributes and without skeleton files
    And user "Alice" has uploaded file with content "some data" to "/textfile0.txt"
    And user "Alice" has created folder "/PARENT"
    And user "Alice" has created folder "/FOLDER"
    And user "Alice" has uploaded file with content "some data" to "/PARENT/parent.txt"

  @smokeTest @skipOnBruteForceProtection @issue-brute_force_protection-112
  Scenario: send MKCOL requests to webDav endpoints as normal user with wrong password
    When user "Alice" requests these endpoints with "MKCOL" including body "doesnotmatter" using password "invalid" about user "Alice"
      | endpoint                                           |
      | /remote.php/webdav/textfile0.txt                   |
      | /remote.php/dav/files/%username%/textfile0.txt     |
      | /remote.php/webdav/PARENT                          |
      | /remote.php/dav/files/%username%/PARENT            |
      | /remote.php/dav/files/%username%/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "401"

  @smokeTest @skipOnBruteForceProtection @issue-brute_force_protection-112 @skipOnOcV10 @personalSpace
  Scenario: send MKCOL requests to webDav endpoints as normal user with wrong password using the spaces WebDAV API
    When user "Alice" requests these endpoints with "MKCOL" including body "doesnotmatter" using password "invalid" about user "Alice"
      | endpoint                                           |
      | /remote.php/dav/spaces/%spaceid%/textfile0.txt     |
      | /remote.php/dav/spaces/%spaceid%/PARENT            |
      | /remote.php/dav/spaces/%spaceid%/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "401"

  @smokeTest @skipOnBruteForceProtection @issue-brute_force_protection-112
  Scenario: send MKCOL requests to webDav endpoints as normal user with no password
    When user "Alice" requests these endpoints with "MKCOL" including body "doesnotmatter" using password "" about user "Alice"
      | endpoint                                           |
      | /remote.php/webdav/textfile0.txt                   |
      | /remote.php/dav/files/%username%/textfile0.txt     |
      | /remote.php/webdav/PARENT                          |
      | /remote.php/dav/files/%username%/PARENT            |
      | /remote.php/dav/files/%username%/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "401"

  @smokeTest @skipOnOcV10 @personalSpace @skipOnBruteForceProtection @issue-brute_force_protection-112
  Scenario: send MKCOL requests to webDav endpoints as normal user with no password using the spaces WebDAV API
    When user "Alice" requests these endpoints with "MKCOL" including body "doesnotmatter" using password "" about user "Alice"
      | endpoint                                           |
      | /remote.php/dav/spaces/%spaceid%/textfile0.txt     |
      | /remote.php/dav/spaces/%spaceid%/PARENT            |
      | /remote.php/dav/spaces/%spaceid%/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "401"

  @skipOnOcV10 @issue-ocis-5049 @issue-ocis-reva-9 @issue-ocis-reva-197
  Scenario: send MKCOL requests to another user's webDav endpoints as normal user
    Given user "Brian" has been created with default attributes and without skeleton files
    When user "Brian" requests these endpoints with "MKCOL" including body "" about user "Alice"
      | endpoint                                           |
      | /remote.php/dav/files/%username%/textfile0.txt     |
      | /remote.php/dav/files/%username%/PARENT            |
      | /remote.php/dav/files/%username%/does-not-exist    |
      | /remote.php/dav/files/%username%/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "404"

  @skipOnOcV10 @issue-ocis-5049 @issue-ocis-reva-9 @issue-ocis-reva-197
  Scenario: send MKCOL requests to non-existent user's webDav endpoints as normal user
    Given user "Brian" has been created with default attributes and without skeleton files
    When user "Brian" requests these endpoints with "MKCOL" including body "" about user "non-existent-user"
      | endpoint                                                  |
      | /remote.php/dav/files/non-existent-user/textfile0.txt     |
      | /remote.php/dav/files/non-existent-user/PARENT            |
      | /remote.php/dav/files/non-existent-user/does-not-exist    |
      | /remote.php/dav/files/non-existent-user/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "404"

  @skipOnOcV10 @personalSpace @issue-ocis-reva-9 @issue-ocis-reva-197
  Scenario: send MKCOL requests to another user's webDav endpoints as normal user using the spaces WebDAV API
    Given user "Brian" has been created with default attributes and without skeleton files
    When user "Brian" requests these endpoints with "MKCOL" including body "" about user "Alice"
      | endpoint                                           |
      | /remote.php/dav/spaces/%spaceid%/textfile0.txt     |
      | /remote.php/dav/spaces/%spaceid%/PARENT            |
      | /remote.php/dav/spaces/%spaceid%/does-not-exist    |
      | /remote.php/dav/spaces/%spaceid%/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "404"

  @skipOnOcV10 @issue-ocis-5049 @personalSpace @issue-ocis-reva-9 @issue-ocis-reva-197
  Scenario: send MKCOL requests to non-existent user's webDav endpoints as normal user using the spaces WebDAV API
    Given user "Brian" has been created with default attributes and without skeleton files
    When user "Brian" requests these endpoints with "MKCOL" including body "" about user "non-existent-user"
      | endpoint                                           |
      | /remote.php/dav/spaces/%spaceid%/textfile0.txt     |
      | /remote.php/dav/spaces/%spaceid%/PARENT            |
      | /remote.php/dav/spaces/%spaceid%/does-not-exist    |
      | /remote.php/dav/spaces/%spaceid%/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "404"


  Scenario: send MKCOL requests to webDav endpoints using invalid username but correct password
    When user "usero" requests these endpoints with "MKCOL" including body "doesnotmatter" using the password of user "Alice"
      | endpoint                                           |
      | /remote.php/webdav/textfile0.txt                   |
      | /remote.php/dav/files/%username%/textfile0.txt     |
      | /remote.php/webdav/PARENT                          |
      | /remote.php/dav/files/%username%/PARENT            |
      | /remote.php/dav/files/%username%/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "401"

  @skipOnOcV10 @personalSpace
  Scenario: send MKCOL requests to webDav endpoints using invalid username but correct password using the spaces WebDAV API
    When user "usero" requests these endpoints with "MKCOL" including body "doesnotmatter" using the password of user "Alice"
      | endpoint                                           |
      | /remote.php/dav/spaces/%spaceid%/textfile0.txt     |
      | /remote.php/dav/spaces/%spaceid%/PARENT            |
      | /remote.php/dav/spaces/%spaceid%/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "401"


  Scenario: send MKCOL requests to webDav endpoints using valid password and username of different user
    Given user "Brian" has been created with default attributes and without skeleton files
    When user "Brian" requests these endpoints with "MKCOL" including body "doesnotmatter" using the password of user "Alice"
      | endpoint                                           |
      | /remote.php/webdav/textfile0.txt                   |
      | /remote.php/dav/files/%username%/textfile0.txt     |
      | /remote.php/webdav/PARENT                          |
      | /remote.php/dav/files/%username%/PARENT            |
      | /remote.php/dav/files/%username%/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "401"

  @skipOnOcV10 @personalSpace
  Scenario: send MKCOL requests to webDav endpoints using valid password and username of different user using the spaces WebDAV API
    Given user "Brian" has been created with default attributes and without skeleton files
    When user "Brian" requests these endpoints with "MKCOL" including body "doesnotmatter" using the password of user "Alice"
      | endpoint                                           |
      | /remote.php/dav/spaces/%spaceid%/textfile0.txt     |
      | /remote.php/dav/spaces/%spaceid%/PARENT            |
      | /remote.php/dav/spaces/%spaceid%/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "401"

  @smokeTest @skipOnBruteForceProtection @issue-brute_force_protection-112
  Scenario: send MKCOL requests to webDav endpoints without any authentication
    When a user requests these endpoints with "MKCOL" with body "doesnotmatter" and no authentication about user "Alice"
      | endpoint                                           |
      | /remote.php/webdav/textfile0.txt                   |
      | /remote.php/dav/files/%username%/textfile0.txt     |
      | /remote.php/webdav/PARENT                          |
      | /remote.php/dav/files/%username%/PARENT            |
      | /remote.php/dav/files/%username%/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "401"

  @smokeTest @skipOnBruteForceProtection @issue-brute_force_protection-112 @skipOnOcV10 @personalSpace
  Scenario: send MKCOL requests to webDav endpoints without any authentication using the spaces WebDAV API
    When a user requests these endpoints with "MKCOL" with body "doesnotmatter" and no authentication about user "Alice"
      | endpoint                                           |
      | /remote.php/dav/spaces/%spaceid%/textfile0.txt     |
      | /remote.php/dav/spaces/%spaceid%/PARENT            |
      | /remote.php/dav/spaces/%spaceid%/PARENT/parent.txt |
    Then the HTTP status code of responses on all endpoints should be "401"
