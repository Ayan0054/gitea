// audit/models.go

package audit

import (
    "time"
)

type AuditEntry struct {
    ID          int64     `xorm:"pk autoincr"`
    UserID      int64     `xorm:"INDEX"`
    RepoID      int64     `xorm:"INDEX"`
    Action      string    // e.g. "create_repo", "delete_repo", "push"
    Description string    // human readable
    Timestamp   time.Time `xorm:"created"`
}
