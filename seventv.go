curl 'https://7tv.io/v3/gql' --compressed -X POST 
-H 'content-type: application/json' 

{"operationName":"SearchEmotes","variables":{"query":"h3llomar","limit":25,"page":1,"sort":{"value":"popularity","order":"DESCENDING"},"filter":{"category":"TOP","exact_match":false,"case_sensitive":false,"ignore_tags":false,"zero_width":false,"animated":false,"aspect_ratio":""}},"query":"query SearchEmotes($query: String!, $page: Int, $sort: Sort, $limit: Int, $filter: EmoteSearchFilter) {\n  emotes(query: $query, page: $page, sort: $sort, limit: $limit, filter: $filter) {\n    count\n    items {\n      id\n      name\n      state\n      trending\n      owner {\n        id\n        username\n        display_name\n        style {\n          color\n          paint_id\n          __typename\n        }\n        __typename\n      }\n      flags\n      host {\n        url\n        files {\n          name\n          format\n          width\n          height\n          __typename\n        }\n        __typename\n      }\n      __typename\n    }\n    __typename\n  }\n}"}