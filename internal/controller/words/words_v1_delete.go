package words

import (
	"context"

	"star/api/words/v1"
)

func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error) {
	uid, err := c.users.GetUid(ctx)
	if err != nil {
		return
	}
	err = c.words.Delete(ctx, uid, req.Id)
	return
}
