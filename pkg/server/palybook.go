package server

import (
    "github.com/kataras/iris/context"
    "kobe/pkg/model/playbook"
)

func init() {
    var cache = playbook.Cache
    url := "/playbooks/"
    App.Get(url, func(ctx context.Context) {
        items := cache.List()
        _, _ = ctx.JSON(items)
    })
    App.Post(url, func(c context.Context) {
        var p playbook.Playbook
        _ = c.ReadJSON(&p)
        cache.Create(p.Name, p)
    })
}
