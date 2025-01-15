import { IconZzz } from '@tabler/icons-react'
import { Handle, NodeProps, Position } from '@xyflow/react'
import { FC } from 'react'

const NodeSleep: FC<NodeProps> = () => {
  return (
    <>
      <Handle type="target" position={Position.Left} />
      <div className="bg-[#7e5cad] p-[15px] rounded-[15px] text-white" style={{ color: 'white' }}>
        <IconZzz />
      </div>
      <Handle type="source" position={Position.Right} />
      <div className="absolute w-full text-center font-bold">Sleep</div>
    </>
  )
}

export default NodeSleep
