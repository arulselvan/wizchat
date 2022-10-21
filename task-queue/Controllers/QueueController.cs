using System.Collections;
using Microsoft.AspNetCore.Mvc;

namespace task_queue.Controllers;

[ApiController]
[Route("[controller]")]

public class QueueController : ControllerBase
{

    private readonly ILogger<QueueController> _logger;

    public QueueController(ILogger<QueueController> logger)
    {
        _logger = logger;
    }

    [HttpPost(Name = "enqueue")]
    public Response Post(Task taskInfo)
    {
        _logger.LogInformation("Inside Enqueue");
        _logger.LogInformation("Task Info:{@task}", taskInfo);
        TaskQueue.taskQueueList.Enqueue(taskInfo);
        return new Response { Status = true, Message = "Task Successfully Enqueued" };
    }

}
