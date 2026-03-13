package pagination

import (
  "net/http"
  "strconv"
)

type Params struct {
  Limit  int
  Offset int
}

func Parse(r *http.Request, defLimit int, maxLimit int) Params {
  q := r.URL.Query()

  limit := defLimit
  if v := q.Get("limit"); v != "" {
    if n, err := strconv.Atoi(v); err == nil {
      limit = n
    }
  }
  if limit <= 0 {
    limit = defLimit
  }
  if limit > maxLimit {
    limit = maxLimit
  }

  offset := 0
  if v := q.Get("offset"); v != "" {
    if n, err := strconv.Atoi(v); err == nil {
      offset = n
    }
  }
  if offset < 0 {
    offset = 0
  }

  return Params{Limit: limit, Offset: offset}
}