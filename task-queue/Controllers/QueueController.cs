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
    public string Post(Task taskInfo)
    {
        _logger.LogInformation("Inside Enqueue");
        _logger.LogInformation("Task Info:{@task}", taskInfo);
        return "Task Successfully Enqueued";
    }
}
