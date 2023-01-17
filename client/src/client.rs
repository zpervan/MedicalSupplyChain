mod medical_supplies_grpc;
mod medical_supplies;

use medical_supplies::*;
use medical_supplies_grpc::*;

use grpc::ClientStubExt;
use futures::executor;

fn main()
{
    let client = MedicalSuppliesClient::new_plain("127.0.0.1", 50000, Default::default()).unwrap();
    let mut req = Request::new();
    req.set_Data("Zvonimir".to_owned());

    let resp = client
        .fetch_all(grpc::RequestOptions::new(), req)
        .join_metadata_result();

    println!("{:?}", executor::block_on(resp));
}