import { DashboardServiceClient } from '@bayu-aditya/ideagate-model-js/dashboard/service.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'

export async function getListApps() {
  const transport = new GrpcWebFetchTransport({
    baseUrl: 'http://localhost:50052',
    format: 'text',
  })

  const dashboardService = new DashboardServiceClient(transport)

  const resp = await dashboardService.getListApplication({ projectId: '1' })
  console.log(resp)
  console.log(resp.response.applications)
}
