.[] | explode | reduce .[] as $char ([]; [$char] + .) | implode
