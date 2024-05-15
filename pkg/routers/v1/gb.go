package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	apsv1 "my.domain/guestbook/apis/apps/v1"
)

var (
	listAllUrl    = fmt.Sprintf("/apis/%s/%s/%s", apsv1.SchemeGroupVersion.Group, apsv1.SchemeGroupVersion.Version, apsv1.GbResourceName)
	listByNsUrl   = fmt.Sprintf("/apis/%s/%s/namespaces/:ns/%s", apsv1.SchemeGroupVersion.Group, apsv1.SchemeGroupVersion.Version, apsv1.GbResourceName)
	detailByNsUrl = fmt.Sprintf("/apis/%s/%s/namespaces/:ns/%s/:name", apsv1.SchemeGroupVersion.Group, apsv1.SchemeGroupVersion.Version, apsv1.GbResourceName)
	postByNsUrl   = fmt.Sprintf("/apis/%s/%s/namespaces/:ns/%s", apsv1.SchemeGroupVersion.Group, apsv1.SchemeGroupVersion.Version, apsv1.GbResourceName)
	patchByNsUrl  = fmt.Sprintf("/apis/%s/%s/namespaces/:ns/%s/:name", apsv1.SchemeGroupVersion.Group, apsv1.SchemeGroupVersion.Version, apsv1.GbResourceName)
)

func RegisterRoute(r *gin.Engine) {
	//// 获取所有
	//r.GET(listAllUrl, func(c *gin.Context) {
	//	list, err := store.NewClientStore().
	//		ListByNsOrAll("") //取全部
	//	if err != nil {
	//		status := utils.NotFoundStatus("Ingress列表不存在")
	//		c.AbortWithStatusJSON(404, status)
	//		return
	//	}
	//	c.JSON(200, utils.ConvertToTable(list))
	//})
	//
	//// 根据 namespace 获取
	//r.GET(listByNsUrl, func(c *gin.Context) {
	//	list, err := store.NewClientStore().
	//		ListByNsOrAll(c.Param("ns"))
	//	if err != nil {
	//		status := utils.NotFoundStatus("Ingress列表不存在")
	//		c.AbortWithStatusJSON(404, status)
	//		return
	//	}
	//	c.JSON(200, utils.ConvertToTable(list))
	//})
	//
	//// 获取具体资源 kubectl get mi mi1
	//r.GET(detailByNsUrl, func(c *gin.Context) {
	//	mi, err := store.NewClientStore().GetByNs(c.Param("name"), c.Param("ns"))
	//	if err != nil {
	//		status := utils.NotFoundStatus(fmt.Sprintf("你要的Ingress:%s在%s这个命名空间没找到，去别地儿看看？",
	//			c.Param("name"), c.Param("ns")))
	//		c.AbortWithStatusJSON(404, status)
	//		return
	//	}
	//	//c.JSON(200, utils.ConvertToTable(mi))
	//	c.JSON(200, mi)
	//})
	//
	//// 新增
	//r.POST(postByNsUrl, func(c *gin.Context) {
	//	mi := &apsv1.GuestBook{}
	//	err := c.ShouldBindJSON(mi)
	//	if err != nil {
	//		c.AbortWithStatusJSON(400, utils.ErrorStatus(400, err.Error(), metav1.StatusReasonBadRequest))
	//		return
	//	}
	//	//创建真实的Ingress
	//	err = builders.CreateIngress(mi)
	//	if err != nil {
	//		c.AbortWithStatusJSON(400, utils.ErrorStatus(400, err.Error(), metav1.StatusReasonBadRequest))
	//		return
	//	}
	//	c.JSON(200, mi)
	//})
	//
	////  如果已经存在， 执行 patch 请求
	//r.PATCH(patchByNsUrl, func(c *gin.Context) {
	//	apply := &apsv1.GuestBook{} //
	//	err := c.ShouldBindJSON(&apply)
	//	if err != nil {
	//		c.AbortWithStatusJSON(400, utils.ErrorStatus(400, err.Error(), metav1.StatusReasonBadRequest))
	//		return
	//	}
	//
	//	newMi, err := builders.PatchIngress(apply)
	//	if err != nil {
	//		c.AbortWithStatusJSON(400, utils.ErrorStatus(400, err.Error(), metav1.StatusReasonBadRequest))
	//		return
	//	}
	//	c.JSON(200, newMi)
	//})
}
