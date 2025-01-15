import { Handle, NodeProps, Position } from '@xyflow/react'
import { FC } from 'react'

const NodeSleep: FC<NodeProps> = () => {
  return (
    <>
      <Handle type="target" position={Position.Left} />
      <div className="bg-[#7e5cad] px-[20px] py-[10px] rounded-md text-white" style={{ color: 'white' }}>
        Sleep
      </div>
      <Handle type="source" position={Position.Right} />
    </>
  )
}

export default NodeSleep
