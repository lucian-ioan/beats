// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package aerospike

import (
	"github.com/elastic/beats/v7/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "aerospike", asset.ModuleFieldsPri, AssetAerospike); err != nil {
		panic(err)
	}
}

// AssetAerospike returns asset data.
// This is the base64 encoded zlib format compressed contents of module/aerospike.
func AssetAerospike() string {
	return "eJzUmMFu4zYQhu9+ikEue6n9AD4UWLSXHnZRFL0tFsGYHFlsKFLgDOMK6MMXlKxElijbcbIbhQcfRHHm/6gZcsZreKBmC0jBc20eaAUgRixt4e5z/+xuBaCJVTC1GO+28OsKAOBpHiqvo01LA1lCpi3scQXAJGLcnrfw7Y7Z3n1fARSGrOZta2ANDis6dZ6GNHUyEXysj08y3k9NDc2lX65R0dNMzuSs2W7kjJzS9WMsYyhFWUNOTqbmtFzQk8ZvrTVgQeHNaDanYqhEkyWhyfQ5NVcoGqjqHIAEdIwqLeC80nNqh4opBB+yb/SirXf7mReu0J3G11jtKIAvjl8qSyElChRoLGk4GCkBXScuRzYIQy/3hY9OL4YhEEcrpME4wKQPWn3nMTgqRcw/DeLor4j2DM95yWIq8nGcdj9O8qV9T3o0+ChT2b3kQJiLk7dKzmT+o6fmmKFTcSEU3i0JJ2onGfhh8m8SPIvMvvyGX5F6h2B+6MXY2v/oyZeBuPFafP94nrIsMqDntnwmpJ8LvUejxvF8c8n5u+EHYPEB99SF7QvrTnxEY3FnaVOr3BZ2ylihJX1fWI+5lwofKpQt1BTUtJ6+AiONL4QcA6VdJKiMM1WsQHknZh99ZNAtair5AVXwzIDWtk/5WDD1LcH8WVIEen/QP7vF6Yv5osNSWKMy0rQCkxeQ0vA1ROIF7WbXCOVS9mzw9zBzi69A+Tt57ww8sRw/kbVeYbpKxY9goD380xOvz5BFJr1AsCQLds31TD1Peajud4FQlTSudjqanfeW0L3sAPijAAmRfhl0+yUy9I7gU2n25fqAQmH9LWH8V1HlQ/N9XSv5dDHUevHdorc6tb601k7PrZc2zAtM5W6TLibzS+Jfo+C7JcHnykcnAzKvVKxNF/9J2SvhjNP07zLp0h3UynslIi+YkUl5pzE0HSnxK1kXdRUdeW8+rdNv9rh7oObgw/gEvyDw65Pnid3nTljTpvSc/zfyNq9eE0xMnjh8Y8zkcJbQ7/4hNSlQb75E/iLlg55p1i7dHhWyUK5LOxumV4Thc5sQjvqGsQaH0qgSMKQaVswjHYVkOpyTAu+nCDVunCbP58FsnrD4+r5thfIf9jVFzViOYVAxBHJim/bvmVRkHozbd70Yb1b/BwAA//89umEy"
}
