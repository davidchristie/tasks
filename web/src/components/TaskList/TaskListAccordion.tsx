import {
  Accordion,
  AccordionDetails,
  AccordionSummary,
  List,
  Typography,
} from "@material-ui/core";
import { makeStyles } from "@material-ui/core/styles";
import { ExpandMore as ExpandMoreIcon } from "@material-ui/icons";
import { useEffect, useState } from "react";
import { Task } from "../../generated/graphql";
import TaskListItem from "../TaskListItem";

interface Props {
  defaultExpanded?: boolean;
  label: string;
  loading: boolean;
  tasks: Pick<Task, "done" | "id" | "text">[];
}

const useStyles = makeStyles({
  list: {
    width: "100%",
  },
});

export default function TaskListAccordion({
  defaultExpanded,
  label,
  loading,
  tasks,
}: Props) {
  const classes = useStyles();

  const [expanded, setExpanded] = useState(defaultExpanded || false);

  const disabled = loading || tasks.length === 0;

  useEffect(() => {
    if (!loading && tasks.length === 0) {
      setExpanded(false);
    }
  }, [loading, tasks.length]);

  const handleChange = (
    event: React.ChangeEvent<{}>,
    nextExpanded: boolean
  ) => {
    setExpanded(nextExpanded);
  };

  return (
    <Accordion
      disabled={disabled}
      expanded={expanded && tasks.length > 0}
      onChange={handleChange}
    >
      <AccordionSummary expandIcon={<ExpandMoreIcon />}>
        <Typography>
          {label} {!loading && <>({tasks.length})</>}
        </Typography>
      </AccordionSummary>
      <AccordionDetails>
        <List className={classes.list}>
          {tasks.map((task, index) => (
            <TaskListItem
              divider={index < tasks.length - 1}
              key={task.id}
              task={task}
            />
          ))}
        </List>
      </AccordionDetails>
    </Accordion>
  );
}
