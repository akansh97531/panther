
<!-- This document is generated by "mage doc:logs". DO NOT EDIT! -->

# OSSEC
##OSSEC.EventInfo
OSSEC EventInfo alert parser. Currently only JSON output is supported.
Reference: https://www.ossec.net/docs/docs/formats/alerts.html
<table>
<tr><th align=center>Column</th><th align=center>Type</th><th align=center>Required</th><th align=center>Description</th></tr>
<tr><td valign=top>id</td><td>string</td><td valign=top>true</td><td valign=top>Unique id of the event.</td></tr>
<tr><td valign=top>rule</td><td>{
<br>&nbsp;&nbsp;&nbsp;&nbsp;"comment": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"group": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"level": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "integer"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"sidid": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "integer"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"CIS": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"items": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "array"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"cve": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"firedtimes": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "integer"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"frequency": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "integer"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"groups": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"items": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "array"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"info": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"PCI_DSS": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"items": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "array"
<br>&nbsp;&nbsp;&nbsp;&nbsp;}
<br>}<br><br></td><td valign=top>true</td><td valign=top>Information about the rule that created the event.</td></tr>
<tr><td valign=top>TimeStamp</td><td>timestamp</td><td valign=top>true</td><td valign=top>Timestamp in UTC.</td></tr>
<tr><td valign=top>location</td><td>string</td><td valign=top>true</td><td valign=top>Source of the event (filename, command, etc).</td></tr>
<tr><td valign=top>hostname</td><td>string</td><td valign=top>true</td><td valign=top>Hostname of the host that created the event.</td></tr>
<tr><td valign=top>full_log</td><td>string</td><td valign=top>true</td><td valign=top>The full captured log of the event.</td></tr>
<tr><td valign=top>action</td><td>string</td><td valign=top>false</td><td valign=top>The event action (drop, deny, accept, etc).</td></tr>
<tr><td valign=top>agentip</td><td>string</td><td valign=top>false</td><td valign=top>The IP address of an agent extracted from the hostname.</td></tr>
<tr><td valign=top>agent_name</td><td>string</td><td valign=top>false</td><td valign=top>The name of an agent extracted from the hostname.</td></tr>
<tr><td valign=top>command</td><td>string</td><td valign=top>false</td><td valign=top>The command extracted by the decoder.</td></tr>
<tr><td valign=top>data</td><td>string</td><td valign=top>false</td><td valign=top>Additional data extracted by the decoder. For example a filename.</td></tr>
<tr><td valign=top>decoder</td><td>string</td><td valign=top>false</td><td valign=top>The name of the decoder used to parse the logs.</td></tr>
<tr><td valign=top>decoder_desc</td><td>{
<br>&nbsp;&nbsp;&nbsp;&nbsp;"accumulate": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "integer"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"fts": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "integer"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"ftscomment": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"name": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"parent": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;}
<br>}<br><br></td><td valign=top>false</td><td valign=top>Information about the decoder used to parse the logs.</td></tr>
<tr><td valign=top>decoder_parent</td><td>string</td><td valign=top>false</td><td valign=top>In the case of a nested decoder, the name of it&#39;s parent.</td></tr>
<tr><td valign=top>dstgeoip</td><td>string</td><td valign=top>false</td><td valign=top>GeoIP location information about the destination IP address.</td></tr>
<tr><td valign=top>dstip</td><td>string</td><td valign=top>false</td><td valign=top>The destination IP address.</td></tr>
<tr><td valign=top>dstport</td><td>string</td><td valign=top>false</td><td valign=top>The destination port.</td></tr>
<tr><td valign=top>dstuser</td><td>string</td><td valign=top>false</td><td valign=top>The destination (target) username.</td></tr>
<tr><td valign=top>logfile</td><td>string</td><td valign=top>false</td><td valign=top>The source log file that was decoded to generate the event.</td></tr>
<tr><td valign=top>previous_output</td><td>string</td><td valign=top>false</td><td valign=top>The full captured log of the previous event.</td></tr>
<tr><td valign=top>program_name</td><td>string</td><td valign=top>false</td><td valign=top>The executable name extracted from the log by the decoder used to match a rule.</td></tr>
<tr><td valign=top>protocol</td><td>string</td><td valign=top>false</td><td valign=top>The protocol (ip, tcp, udp, etc) extracted by the decoder.</td></tr>
<tr><td valign=top>srcgeoip</td><td>string</td><td valign=top>false</td><td valign=top>GeoIP location information about the source IP address.</td></tr>
<tr><td valign=top>srcip</td><td>string</td><td valign=top>false</td><td valign=top>The source IP address.</td></tr>
<tr><td valign=top>srcport</td><td>string</td><td valign=top>false</td><td valign=top>The source port.</td></tr>
<tr><td valign=top>srcuser</td><td>string</td><td valign=top>false</td><td valign=top>The source username.</td></tr>
<tr><td valign=top>status</td><td>string</td><td valign=top>false</td><td valign=top>Event status (success, failure, etc).</td></tr>
<tr><td valign=top>SyscheckFile</td><td>{
<br>&nbsp;&nbsp;&nbsp;&nbsp;"gowner_after": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"gowner_before": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"md5_after": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"md5_before": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"owner_after": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"owner_before": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"path": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"perm_after": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "integer"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"perm_before": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "integer"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"sha1_after": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;"sha1_before": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;}
<br>}<br><br></td><td valign=top>false</td><td valign=top>Information about a file integrity check.</td></tr>
<tr><td valign=top>systemname</td><td>string</td><td valign=top>false</td><td valign=top>The system name extracted by the decoder.</td></tr>
<tr><td valign=top>url</td><td>string</td><td valign=top>false</td><td valign=top>URL of the event.</td></tr>
<tr><td valign=top>p_log_type</td><td>string</td><td valign=top>true</td><td valign=top>Panther added field with type of log</td></tr>
<tr><td valign=top>p_row_id</td><td>string</td><td valign=top>true</td><td valign=top>Panther added field with unique id (within table)</td></tr>
<tr><td valign=top>p_event_time</td><td>timestamp</td><td valign=top>true</td><td valign=top>Panther added standardize event time (UTC)</td></tr>
<tr><td valign=top>p_parse_time</td><td>timestamp</td><td valign=top>true</td><td valign=top>Panther added standardize log parse time (UTC)</td></tr>
<tr><td valign=top>p_any_ip_addresses</td><td>{
<br>&nbsp;&nbsp;&nbsp;&nbsp;"p_any_ip_addresses": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"items": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "array"
<br>&nbsp;&nbsp;&nbsp;&nbsp;}
<br>}<br><br></td><td valign=top>false</td><td valign=top>Panther added field with collection of ip addresses associated with the row</td></tr>
<tr><td valign=top>p_any_domain_names</td><td>{
<br>&nbsp;&nbsp;&nbsp;&nbsp;"p_any_domain_names": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"items": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "array"
<br>&nbsp;&nbsp;&nbsp;&nbsp;}
<br>}<br><br></td><td valign=top>false</td><td valign=top>Panther added field with collection of domain names associated with the row</td></tr>
<tr><td valign=top>p_any_sha1_hashes</td><td>{
<br>&nbsp;&nbsp;&nbsp;&nbsp;"p_any_sha1_hashes": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"items": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "array"
<br>&nbsp;&nbsp;&nbsp;&nbsp;}
<br>}<br><br></td><td valign=top>false</td><td valign=top>Panther added field with collection of SHA1 hashes associated with the row</td></tr>
<tr><td valign=top>p_any_md5_hashes</td><td>{
<br>&nbsp;&nbsp;&nbsp;&nbsp;"p_any_md5_hashes": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"items": {
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "string"
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;},
<br>&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;"type": "array"
<br>&nbsp;&nbsp;&nbsp;&nbsp;}
<br>}<br><br></td><td valign=top>false</td><td valign=top>Panther added field with collection of MD5 hashes associated with the row</td></tr>
</table>
