mod proto;

use proto::{medical_supplies::*, medical_supplies_grpc::*};

use grpc::ClientStubExt;
use futures::executor;

fn main()
{
    let client = MedicalSuppliesClient::new_plain("127.0.0.1", 3500, Default::default()).unwrap();
    let mut req = Request::new();
    req.set_Data("Zvonimir".to_owned());

    let resp = client
        .test_request(grpc::RequestOptions::new(), req)
        .join_metadata_result();

    println!("{:?}", executor::block_on(resp));
}