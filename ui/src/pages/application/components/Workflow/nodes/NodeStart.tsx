import { IconPlayerPlay } from '@tabler/icons-react'
import { Handle, NodeProps, Position } from '@xyflow/react'
import { FC } from 'react'

const NodeStart: FC<NodeProps> = () => {
  return (
    <>
      <div className="bg-[#474e93] p-[15px] rounded-[15px] text-white" style={{ color: 'white' }}>
        <IconPlayerPlay />
      </div>
      <Handle type="source" position={Position.Right} />
      <div className="absolute w-full text-center font-bold">Start</div>
    </>
  )
}

export default NodeStart
