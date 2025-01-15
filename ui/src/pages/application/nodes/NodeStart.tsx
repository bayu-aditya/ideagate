import { Handle, NodeProps, Position } from '@xyflow/react'
import { FC } from 'react'

const NodeStart: FC<NodeProps> = () => {
  return (
    <>
      <div className="bg-[#474e93] px-[20px] py-[10px] rounded-md text-white" style={{ color: 'white' }}>
        Start
      </div>
      <Handle type="source" position={Position.Right} />
    </>
  )
}

export default NodeStart
