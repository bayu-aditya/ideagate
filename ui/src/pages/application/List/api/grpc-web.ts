import { DashboardServiceClient } from '@bayu-aditya/ideagate-model-js/dashboard/service.client'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport'

export async function getListApps(projectId: string) {
  const transport = new GrpcWebFetchTransport({
    baseUrl: 'http://localhost:50052',
    format: 'text',
  })

  const dashboardService = new DashboardServiceClient(transport)

  const resp = await dashboardService.getListApplication({ projectId })

  return resp.response
}
