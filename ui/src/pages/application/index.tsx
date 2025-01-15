import '@xyflow/react/dist/style.css'

import {
  addEdge,
  Background,
  BackgroundVariant,
  Connection,
  Controls,
  MiniMap,
  ReactFlow,
  useEdgesState,
  useNodesState,
} from '@xyflow/react'
import { FC, useCallback } from 'react'

import { nodeTypes } from './nodes'

const initialNodes = [
  { id: '1', type: 'start', position: { x: 0, y: 0 }, data: { label: '1' } },
  { id: '2', type: 'sleep', position: { x: 150, y: 0 }, data: { label: '2' } },
  { id: '3', type: 'sleep', position: { x: 300, y: 0 }, data: { label: '3' } },
  { id: '4', type: 'sleep', position: { x: 230, y: 100 }, data: { label: '4' } },
  { id: '5', type: 'end', position: { x: 450, y: 0 }, data: { label: '5' } },
]
const initialEdges = [{ id: 'e1-2', source: '1', target: '2', animated: true }]

const ApplicationPage: FC = () => {
  const [nodes, , onNodesChange] = useNodesState(initialNodes)
  const [edges, setEdges, onEdgesChange] = useEdgesState(initialEdges)

  const onConnect = useCallback(
    (params: Connection) => setEdges((eds) => addEdge({ ...params, animated: true }, eds)),
    [setEdges]
  )

  return (
    <div>
      <div style={{ width: '100%', height: '600px', backgroundColor: 'white' }}>
        <ReactFlow
          nodes={nodes}
          edges={edges}
          onNodesChange={onNodesChange}
          onEdgesChange={onEdgesChange}
          onConnect={onConnect}
          proOptions={{ hideAttribution: true }}
          nodeTypes={nodeTypes}
          fitView
        >
          <Controls />
          <MiniMap />
          <Background variant={BackgroundVariant.Dots} gap={12} size={1} />
        </ReactFlow>
      </div>
    </div>
  )
}

export default ApplicationPage
