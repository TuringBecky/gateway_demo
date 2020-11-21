package tcp_proxy_middleware

import (
	"fmt"
	"github.com/e421083458/go_gateway/dao"
	"github.com/e421083458/go_gateway/public"
)

func TcpServerFlowCountMiddleware() func(c *TcpSliceRouterContext) {
	return func(c *TcpSliceRouterContext) {
		tmp := c.Get("service_detail")
		fmt.Println("tmp", tmp)
		if tmp == nil {
			c.conn.Write([]byte("TcpServerFlowCountMiddleware get service_detail error"))
			c.Abort()
			return
		}
		serviceDetail := tmp.(*dao.ServiceDetail)

		totalCounter, err := public.FlowCounterHandler.GetCounter(public.FlowTotal)
		if err != nil {
			c.conn.Write([]byte(err.Error()))
			c.Abort()
			return
		}
		totalCounter.Increase()

		counter, err := public.FlowCounterHandler.GetCounter(public.FlowServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			c.conn.Write([]byte(err.Error()))
			c.Abort()
			return
		}
		counter.Increase()
		c.Next()
	}
}
