import React from 'react'
import styles from './Resources.module.css'
import { OutputJSONType } from '../../../consts/types'
import ResourceGraph from './ResourceGraph'

const Resources = (props: { output: OutputJSONType }) => {
  return (
    <div className={styles.root}>
      <h1>Resources over Time</h1>
      <p>
        Select teams to show/hide by clicking the team at the top of the chart.
        Use the slider at the bottom to change which turns are displayed.
      </p>
      <ResourceGraph output={props.output} />
    </div>
  )
}

export default Resources
