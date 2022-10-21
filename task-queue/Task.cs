using System.Collections;

namespace task_queue;

public class Task
{
    public string? TaskType { get; set; }
    public string? UserId { get; set; }
    public string? QueueName { get; set; }
    public string? Priority { get; set; }
}

public class Response
{
    public bool Status { get; set; }
    public string? Message { get; set; }
}

public static class TaskQueue
{
    public static Queue<Task> taskQueueList = new Queue<Task>();
}

