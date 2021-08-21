#!/usr/bin/awk -f

# Copyright 2021 Clément Joly
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
#     Unless required by applicable law or agreed to in writing, software
#     distributed under the License is distributed on an "AS IS" BASIS,
#     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#     See the License for the specific language governing permissions and
#     limitations under the License.
BEGIN {
	removing = 0	# Sequence to remove started
	inserting = 0	# Sequence to insert the content of the comment started
}

/^<!--+ remove -+->$/ {
	removing = 1
	next
}

/^<!--+ end_remove -+->$/ {
	removing = 0
	next
}

/^<!--+ insert$/ {
	inserting = 1
	next
}

/^end_insert -+->$/ {
	inserting = 0
	next
}

removing == 0 {
	print $0
}

